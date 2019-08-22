package main

import (
	"fmt"
	"github.com/xshoji/go-sample-box/grpc-gateway-rest/gateway"
	"github.com/xshoji/go-sample-box/grpc-gateway-rest/proto"
	"github.com/xshoji/go-sample-box/grpc-gateway-rest/proto/impl"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	GRPC_ENDPOINT    = "localhost:19090"
	GATEWAY_ENDPOINT = "localhost:8080"
)

func main() {

	finish := make(chan bool)

	go func() {
		lis, err := net.Listen("tcp", GRPC_ENDPOINT)
		if err != nil {
			log.Fatal(err)
		}

		s := grpc.NewServer()
		proto.RegisterUserServiceServer(s, new(proto_impl.UserService))
		fmt.Printf("Start grpc server: %v\n", GRPC_ENDPOINT)
		err = s.Serve(lis)
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		fmt.Printf("Start gateway server: %v\n", GATEWAY_ENDPOINT)
		if err := gateway.Run(GATEWAY_ENDPOINT, GRPC_ENDPOINT); err != nil {
			panic(err)
		}
	}()

	<-finish

}
