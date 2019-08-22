## About

grpc-gateway sample

## Preparation

```
brew install protobuf
go get -u google.golang.org/grpc
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```

## Build and run

```
// generage code
bash generator.sh

// run grpc and gateway server
go run main.go
```

## How to implement

### Define proto file

Create sample.proto in `/proto`.

> Language Guide (proto3)  ｜  Protocol Buffers  ｜  Google Developers  
> https://developers.google.com/protocol-buffers/docs/proto3

### Generate golang codes

```
bash generator.sh
```

### Implement service

Implement sample.pb.go interfaces in `/proto/impl`.

```
package proto_impl
...
type SampleService struct{}

func (*SampleService) Create(context context.Context, user *proto.User) (*proto.SampleServiceResponse, error) {...}
func (*SampleService) Read(ctx context.Context, in *proto.SampleServiceSelector) (*proto.SampleServiceResponse, error) {...}
func (*SampleService) ReadAll(ctx context.Context, in *empty.Empty) (*proto.SampleServiceResponse, error) {...}
func (*SampleService) Update(context context.Context, user *proto.User) (*proto.SampleServiceResponse, error) {...}
func (*SampleService) Delete(context context.Context, in *proto.SampleServiceSelector) (*proto.SampleServiceResponse, error) {...}

```

### Register handler to gateway.go

```
// newGateway()
...
	err = proto.RegisterSampleServiceHandler(ctx, mux, conn)
...
```

### Register service to main.go

```
// main()
...
		proto.RegisterSampleServiceServer(s, new(proto_impl.UserService))
...
```

## References

> grpc-gatewayでgRPCサーバをRESTで叩けるようにする - Carpe Diem  
> https://christina04.hatenablog.com/entry/2017/11/15/034455

> gRPCを使ってみた - チャーリー！のテクメモ  
> http://charleysdiary.hatenablog.com/entry/2016/09/08/163909

> Protocol Buffers(proto3)でoptionalをどう扱うか - Qiita  
> https://qiita.com/disc99/items/a8ac2a264f322bc6d6e5

> Usage ｜ grpc-gateway  
> https://grpc-ecosystem.github.io/grpc-gateway/docs/usage.html