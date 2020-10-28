module go.xixo.com/api

go 1.15

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-playground/validator/v10 v10.4.1
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/jackc/pgx/v4 v4.9.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.3.0
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.5.1
	github.com/vektah/gqlparser v1.3.1
	github.com/vektah/gqlparser/v2 v2.1.0
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20201002170205-7f63de1d35b0
	google.golang.org/genproto v0.0.0-20201007142714-5c0e72c5e71e
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.3.0 // indirect
	go.xixo.com/protobuf v1.0.0
)

replace go.xixo.com/protobuf v1.0.0 => ../protobuf
