package authr

import (
	"context"

	"go.xixo.com/api/gateway/auth"
	"go.xixo.com/api/pkg/token"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// ClientInterceptor authentication gRPC interceptor
type ClientInterceptor struct {
	jwtManager *token.JWTManager
	logger     *zap.Logger
	checker    Checker
	whiteList  map[string]bool
}

// NewClientInterceptor return new instance of auth gRPC interceptor
func NewClientInterceptor() *ClientInterceptor {
	return &ClientInterceptor{}
}

// Unary .
func (*ClientInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
	) error {
		token, ok := auth.TokenFromContext(ctx)
		if !ok {
			invoker(ctx, method, req, reply, cc, opts...)
		}
		return invoker(metadata.AppendToOutgoingContext(ctx, "authorization", token), method, req, reply, cc, opts...)
	}
}

// Stream .
func (*ClientInterceptor) Stream() grpc.StreamClientInterceptor {
	panic("Unimplemented")
}
