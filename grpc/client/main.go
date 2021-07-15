package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:9000")
	if err != nil {
		fmt.Println(err)
		return
	}
	var reply string
	err = client.Call("HelloService.HelloWord", "hello", &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reply)
}
