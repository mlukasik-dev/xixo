## Generate code for identity service

```
protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative .\identitypb\identity.proto
```

## Generate code for account service

```
protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative .\accountpb\account.proto
```
