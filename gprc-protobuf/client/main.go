package main

import (
	"context"
	"execlt1/gprc-protobuf/pb"
	"fmt"
	"github.com/lack-io/vine"

)


func main() {
	svc := vine.NewService(vine.Name("tt"))
	service := pb.NewRpcService("tt", svc.Client())

	rsp, err := service.HelloWorld(context.TODO(), &pb.HelloWorldRequest{Name: "world"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("result: %v\n", rsp)
}
