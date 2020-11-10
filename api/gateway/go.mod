module go.xixo.com/api/gateway

go 1.15

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/google/uuid v1.1.2
	github.com/joho/godotenv v1.3.0
	github.com/vektah/gqlparser v1.3.1
	github.com/vektah/gqlparser/v2 v2.1.0
	go.xixo.com/api/pkg v1.0.0 // local
	go.xixo.com/api/services/account v1.0.0 // local
	go.xixo.com/api/services/identity v1.0.0 // local
	go.xixo.com/protobuf v1.0.0 // local
	google.golang.org/genproto v0.0.0-20201106154455-f9bfe239b0ba
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.25.0
)

replace go.xixo.com/protobuf v1.0.0 => ../../protobuf

replace go.xixo.com/api/pkg v1.0.0 => ../pkg

replace go.xixo.com/api/services/identity v1.0.0 => ../services/identity

replace go.xixo.com/api/services/account v1.0.0 => ../services/account
