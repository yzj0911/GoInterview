package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (service *HelloService) HelloWord(req string, replay *string) error {
	*replay = "hello:" + req
	return nil
}

func main() {

	rpc.RegisterName(
		"HelloService",
		new(HelloService),
	)
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	rpc.ServeConn(conn)
}
