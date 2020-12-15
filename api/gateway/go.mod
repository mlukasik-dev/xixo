module go.xixo.com/api/gateway

go 1.15

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/agnivade/levenshtein v1.1.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/uuid v1.1.2
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/mitchellh/mapstructure v1.4.0 // indirect
	github.com/vektah/gqlparser v1.3.1
	github.com/vektah/gqlparser/v2 v2.1.0
	golang.org/x/net v0.0.0-20201202161906-c7110b5ffcbb // indirect
	golang.org/x/sys v0.0.0-20201204225414-ed752295db88 // indirect
	golang.org/x/text v0.3.4 // indirect
	go.xixo.com/api/pkg v0.1.0 // local
	go.xixo.com/api/services/account v0.1.0 // local
	go.xixo.com/api/services/identity v0.1.0 // local
	go.xixo.com/protobuf v0.1.0 // local
	google.golang.org/genproto v0.0.0-20201204160425-06b3db808446
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
)

replace go.xixo.com/protobuf v0.1.0 => ../../protobuf

replace go.xixo.com/api/pkg v0.1.0 => ../pkg

replace go.xixo.com/api/services/identity v0.1.0 => ../services/identity

replace go.xixo.com/api/services/account v0.1.0 => ../services/account
