module go.xixo.com/api/services/identity

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/go-playground/validator/v10 v10.4.1
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/jackc/pgx/v4 v4.9.2
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.3.0
	go.uber.org/zap v1.16.0
	go.xixo.com/api/pkg v1.0.0 // local
	go.xixo.com/protobuf v1.0.0 // local
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	google.golang.org/genproto v0.0.0-20201106154455-f9bfe239b0ba
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.25.0
)

replace go.xixo.com/protobuf v1.0.0 => ../../../protobuf

replace go.xixo.com/api/pkg v1.0.0 => ../../pkg
