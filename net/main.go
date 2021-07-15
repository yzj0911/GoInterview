package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("======", i, "   start==============")
			l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", i))
			if err != nil {
				fmt.Println("11111111")
				fmt.Println(err)
				return
			}
			err = l.Close()
			if err != nil {
				fmt.Println("===========")
				fmt.Println(err)
				return
			}
			fmt.Println("end")
			wg.Done()
		}(i)
	}
	wg.Wait()
}
