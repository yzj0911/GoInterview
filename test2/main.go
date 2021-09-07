package main

import (
	"fmt"
	"sync"
)

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	//for i := 0; i <= 4; i++ {
	//	wg.Add(1)
	//	go add(i)
	//}
	//wg.Wait()
	//a := make(chan int, 0)
	//a <- 1
	//b := <-a
	//fmt.Println(b)


	var m =make(map[int]string,1)
	m[1]="!"
	m[2]="11"
	fmt.Println(m)
}

var lock sync.Mutex
var wg sync.WaitGroup
var s = 1000

func add(count int) {
	lock.Lock()
	fmt.Printf("加锁----第%d个携程\n", count)
	defer func() {
		fmt.Printf("解锁----第%d个携程\n", count)
		lock.Unlock()
	}()
	for i := 0; i < 4; i++ {
		s++
		fmt.Printf("j %d gorount %d \n", s, count)
	}
	wg.Done()
}
