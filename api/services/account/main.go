package main

import (
	"context"
	"log"
	"net"
	"time"

	"go.xixo.com/api/gateway/auth"
	"go.xixo.com/api/pkg/authr"
	"go.xixo.com/api/pkg/token"
	"go.xixo.com/api/services/account/domain/accounts"
	"go.xixo.com/api/services/account/grpc/controller"
	"go.xixo.com/api/services/account/postgres"
	"go.xixo.com/protobuf/accountpb"
	"go.xixo.com/protobuf/identitypb"

	"github.com/go-playground/validator/v10"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to create logger: %v\n", err)
	}
	defer logger.Sync()

	identityConn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithUnaryInterceptor(unaryAuthInterceptor))
	if err != nil {
		log.Fatalf("Err: %v\n", err)
	}
	usersClient := identitypb.NewUsersClient(identityConn)

	repo := postgres.NewRepository(postgres.MustConnect(), logger)
	jwtManager := token.NewJWTManager("secret", time.Hour*24)
	validate := validator.New()

	accountsSvc := accounts.New(repo, logger, validate)

	accountsCtr := controller.NewAccountsController(accountsSvc, usersClient)

	authIntr := authr.NewServerInterceptor(&authr.ServerInterceptorConfig{
		JWTManager: jwtManager,
		Logger:     logger,
		Checker:    repo,
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
	accountpb.RegisterAccountsServer(s, accountsCtr)

	// TODO: move to config file
	l, err := net.Listen("tcp", ":"+"50052")
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
	token, ok := auth.TokenFromContext(ctx)
	if !ok {
		invoker(ctx, method, req, reply, cc, opts...)
	}
	return invoker(metadata.AppendToOutgoingContext(ctx, "authorization", token), method, req, reply, cc, opts...)
}
