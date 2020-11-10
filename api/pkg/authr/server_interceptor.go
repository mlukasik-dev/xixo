package authr

import (
	"context"
	"fmt"

	"go.xixo.com/api/pkg/token"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Checker .
type Checker interface {
	CheckPermission(ctx context.Context, roleID uuid.UUID, method string) (bool, error)
}

// ServerInterceptor authentication gRPC interceptor
type ServerInterceptor struct {
	jwtManager *token.JWTManager
	logger     *zap.Logger
	checker    Checker
	whiteList  map[string]bool
}

// ServerInterceptorConfig configuration object
type ServerInterceptorConfig struct {
	JWTManager *token.JWTManager
	Logger     *zap.Logger
	Checker    Checker
	WhiteList  map[string]bool
}

// NewServerInterceptor return new instance of auth gRPC interceptor
func NewServerInterceptor(config *ServerInterceptorConfig) *ServerInterceptor {
	whiteList := config.WhiteList
	if config.WhiteList == nil {
		whiteList = make(map[string]bool)
	}
	return &ServerInterceptor{
		jwtManager: config.JWTManager,
		logger:     config.Logger,
		checker:    config.Checker,
		whiteList:  whiteList,
	}
}

type contextKey string

var tokenContextKey contextKey = "token"

// Unary gRPC unary interceptor
func (i *ServerInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		token, err := i.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		h, err := handler(context.WithValue(ctx, tokenContextKey, token), req)
		return h, err
	}
}

// Stream gRPC streaming interceptor
func (i *ServerInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return handler(srv, stream)
	}
}

func (i *ServerInterceptor) authorize(ctx context.Context, method string) (*token.JWTClaims, error) {
	fmt.Printf("white list: %v\nmethod: %s\n", i.whiteList, method)
	if i.whiteList[method] {
		return nil, nil
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}
	values := md["authorization"]
	if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}
	token := values[0]
	claims, err := i.jwtManager.Verify(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}
	roleIDs := claims.RoleIDs
	var hasPermission bool
	for _, roleID := range roleIDs {
		hasPermission, err = i.checker.CheckPermission(context.Background(), roleID, method)
		fmt.Printf("checking for %s and %s, result: %t\n", roleID, method, hasPermission)
		if err != nil {
			return nil, err
		}
		if hasPermission {
			break
		}
	}
	if !hasPermission {
		return nil, status.Error(codes.PermissionDenied, "no permission to access this RPC")
	}
	return claims, nil
}

// ClaimsFromContext retreives jwt token from context
func ClaimsFromContext(ctx context.Context) (*token.JWTClaims, bool) {
	token, ok := ctx.Value(tokenContextKey).(*token.JWTClaims)
	return token, ok
}
