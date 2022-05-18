package main

import (
	"context"
	"execlt1/gprc-protobuf/pb"
	"fmt"

	"github.com/lack-io/vine"
)

//protoc -I. --go_out=plugins=grpc:. pb/hello.proto
type HelloWorld struct {
}

func (t *HelloWorld) HelloWorld(ctx context.Context, request *pb.HelloWorldRequest, response *pb.HelloWorldResponse) error {
	response.Reply = fmt.Sprintf("%s %s ", request.Name, "hello")
	fmt.Println(response.Reply)
	return nil
}

func main() {
	service := vine.NewService(
		vine.Name("tt"),
		vine.Address("127.0.0.1:9000"),
	)

	service.Init()

	pb.RegisterRpcHandler(service.Server(), new(HelloWorld))

	service.Run()
}
