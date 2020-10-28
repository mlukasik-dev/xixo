package main

import (
	"flag"
	"log"
	"net/http"

	"go.xixo.com/api/pkg/authr"

	"go.xixo.com/api/gateway/auth"
	"go.xixo.com/api/gateway/graph"
	"go.xixo.com/api/gateway/graph/generated"
	"go.xixo.com/protobuf/accountpb"
	"go.xixo.com/protobuf/identitypb"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"google.golang.org/grpc"
)

var port = flag.String("p", "8080", "port")

func main() {
	flag.Parse()

	authIntr := authr.NewClientInterceptor()

	identityConn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithUnaryInterceptor(authIntr.Unary()))
	if err != nil {
		log.Fatalf("Err: %v\n", err)
	}
	accountConn, err := grpc.Dial(":50052", grpc.WithInsecure(), grpc.WithUnaryInterceptor(authIntr.Unary()))
	if err != nil {
		log.Fatalf("Err: %v\n", err)
	}

	authClient := identitypb.NewAuthClient(identityConn)
	rolesClient := identitypb.NewRolesClient(identityConn)
	usersClient := identitypb.NewUsersClient(identityConn)
	accountsClient := accountpb.NewAccountsClient(accountConn)

	r := graph.NewResolver(&graph.Clients{
		AuthClient:     authClient,
		RolesClient:    rolesClient,
		UsersClient:    usersClient,
		AccountsClient: accountsClient,
	})
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: r}))

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(auth.Middleware())
	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", *port)
	log.Fatal(http.ListenAndServe(":"+*port, router))
}
