package main

import (
	"log"
	"net"
	"time"

	"go.xixo.com/api/pkg/authr"
	"go.xixo.com/api/pkg/token"
	"go.xixo.com/api/services/identity/domain/admins"
	"go.xixo.com/api/services/identity/domain/auth"
	"go.xixo.com/api/services/identity/domain/roles"
	"go.xixo.com/api/services/identity/domain/users"
	"go.xixo.com/api/services/identity/grpc/controller"
	"go.xixo.com/api/services/identity/postgres"
	"go.xixo.com/protobuf/identitypb"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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

	repo := postgres.NewRepository(postgres.MustConnect())
	jwtManager := token.NewJWTManager("secret", time.Hour*24)
	validate := validator.New()

	authSvc := auth.New(repo, jwtManager, validate)
	adminsSvc := admins.New(repo, validate)
	usersSvc := users.New(repo, validate)
	rolesSvc := roles.New(repo, validate)

	authCtr := controller.New(&controller.Services{
		AuthSvc:   authSvc,
		AdminsSvc: adminsSvc,
		UsersSvc:  usersSvc,
		RolesSvc:  rolesSvc,
	})

	authIntr := authr.NewServerInterceptor(&authr.ServerInterceptorConfig{
		JWTManager: jwtManager,
		Logger:     logger,
		Checker:    repo,
		WhiteList: map[string]bool{
			"/xixo.identity.v1.Auth/Login":    true,
			"/xixo.identity.v1.Auth/Register": true,
		},
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
	identitypb.RegisterIdentityServiceServer(s, authCtr)

	// TODO: move to config file
	l, err := net.Listen("tcp", ":"+"50051")
	if err != nil {
		panic(err)
	}

	log.Printf("Starting gRPC identity service on %s...\n", l.Addr().String())
	if err = s.Serve(l); err != nil {
		log.Fatalf("Failed to start identity service: %v\n", err)
	}
}
