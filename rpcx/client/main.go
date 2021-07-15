package main
//
//import (
//	"context"
//	"execlt1/rpcx/example"
//	"flag"
//	"fmt"
//	"github.com/smallnest/rpcx/client"
//	"log"
//)
//
//var (
//	addr = flag.String("addr", "127.0.0.1:9000", "server address")
//)
//
//func main() {
//	Peer2Peer()
//}
//func Peer2Peer() {
//	flag.Parse()
//	d, err := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
//	if err != nil {
//		fmt.Println(err)
//	}
//	xclinet := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
//	defer xclinet.Close()
//	args := &example.Args{
//		A: 10,
//		B: 20,
//	}
//	reply := &example.Reply{}
//	err = xclinet.Call(context.Background(), "Mul", args, reply)
//	if err != nil {
//		log.Fatalf("failed to call: %v", err)
//	}
//
//}
