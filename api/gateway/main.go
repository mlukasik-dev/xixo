package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"go.xixo.com/api/gateway/auth"
	"go.xixo.com/api/gateway/graph"
	"go.xixo.com/api/gateway/graph/generated"
	"go.xixo.com/api/pkg/authr"
	"go.xixo.com/protobuf/accountpb"
	"go.xixo.com/protobuf/identitypb"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	var envFile string
	flag.StringVar(&envFile, "env-file", "", "env file")
	flag.Parse()

	if envFile != "" {
		err := godotenv.Load(envFile)
		if err != nil {
			log.Fatalf("Failed to load environment variables from: %s\n", envFile)
		} else {
			log.Printf("Loaded environment variables from: %s\n", envFile)
		}
	}
	port := os.Getenv("PORT")
	identitySvcHost := os.Getenv("IDENTITY_SERVICE_HOST")
	identitySvcPort := os.Getenv("IDENTITY_SERVICE_PORT")
	accountSvcHost := os.Getenv("ACCOUNT_SERVICE_HOST")
	accountSvcPort := os.Getenv("ACCOUNT_SERVICE_PORT")
	if port == "" {
		log.Fatalln("Port not provided")
	}

	authIntr := authr.NewClientInterceptor()

	identityConn, err := grpc.Dial(identitySvcHost+":"+identitySvcPort, grpc.WithInsecure(), grpc.WithUnaryInterceptor(authIntr.Unary()))
	if err != nil {
		log.Fatalf("Err: %v\n", err)
	}
	accountConn, err := grpc.Dial(accountSvcHost+":"+accountSvcPort, grpc.WithInsecure(), grpc.WithUnaryInterceptor(authIntr.Unary()))
	if err != nil {
		log.Fatalf("Err: %v\n", err)
	}

	identitySvcClient := identitypb.NewIdentityServiceClient(identityConn)
	accountSvcClient := accountpb.NewAccountServiceClient(accountConn)

	r := graph.NewResolver(&graph.Clients{
		IdentitySvcClient: identitySvcClient,
		AccountSvcClient:  accountSvcClient,
	})
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: r}))

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(auth.Middleware())
	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
