package main

import (
	"context"
	"log"
	"net"

	"go.xixo.com/api/pkg/authr"
	"go.xixo.com/api/pkg/token"
	"go.xixo.com/api/services/account/config"
	"go.xixo.com/api/services/account/domain/accounts"
	"go.xixo.com/api/services/account/grpc/controller"
	"go.xixo.com/api/services/account/postgres"
	"go.xixo.com/protobuf/accountpb"
	"go.xixo.com/protobuf/identitypb"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to create logger: %v\n", err)
	}
	defer logger.Sync()

	identityConn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithUnaryInterceptor(unaryAuthInterceptor))
	if err != nil {
		log.Fatalf("Failed to dial identity service: %v\n", err)
	}
	identitySvcClient := identitypb.NewIdentityServiceClient(identityConn)

	repo := postgres.NewRepository(postgres.MustConnect(), logger)
	jwtManager := token.NewJWTManager(config.Global.Auth.Secret, config.Global.Auth.TokenDuration)
	validate := validator.New()

	accountsSvc := accounts.New(repo, logger, validate)
	accountsCtr := controller.New(accountsSvc, identitySvcClient)

	authIntr := authr.NewServerInterceptor(&authr.ServerInterceptorConfig{
		JWTManager: jwtManager,
		Logger:     logger,
		Checker:    nil, // TODO: use real value
	})

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(logger),
			authIntr.Stream(),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger),
			authIntr.Unary(),
			// grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	reflection.Register(s)
	accountpb.RegisterAccountServiceServer(s, accountsCtr)

	l, err := net.Listen("tcp", config.Global.Port.String())
	if err != nil {
		panic(err)
	}

	log.Printf("Starting gRPC accounts service on %s...\n", l.Addr().String())
	if err = s.Serve(l); err != nil {
		panic(err)
	}
}

// TODO: move to pkg
func unaryAuthInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	token, ok := token.FromContext(ctx)
	if !ok {
		invoker(ctx, method, req, reply, cc, opts...)
	}
	return invoker(metadata.AppendToOutgoingContext(ctx, "authorization", token), method, req, reply, cc, opts...)
}
