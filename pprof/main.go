package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	ip := "127.0.0.1:9000"
	if err := http.ListenAndServe(ip, nil); err != nil {
		fmt.Printf("start pprof failed on %s\n", ip)
	}
}