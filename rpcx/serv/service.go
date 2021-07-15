package main

import (
	"flag"
	"github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:9000", "server address")
)

func main() {
	flag.Parse()
	s := server.Server{}
	s.RegisterName("Arith", new(example.Arith), "")
	go s.Serve("tcp", *addr)
	select {}
}