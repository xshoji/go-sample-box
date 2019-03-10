## About

grpc-gateway sample

## Preparation

```
brew install protobuf
go get -u google.golang.org/grpc
go get -u go.pedge.io/protoeasy/cmd/protoeasy
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```

## Build and run

```
// generage code
protoeasy --go --grpc --grpc --grpc-gateway .

// run grpc and gateway server
go run main.go
```

## References

> grpc-gatewayでgRPCサーバをRESTで叩けるようにする - Carpe Diem  
> https://christina04.hatenablog.com/entry/2017/11/15/034455