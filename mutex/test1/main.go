package main

import (
	"fmt"
	"sync"
)

var CMap struct {
	C int
	sync.Mutex
}

func work(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		CMap.Lock()
		CMap.C += i
		CMap.Unlock()
	}
	fmt.Println(CMap.C)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go work(&wg)
	go work(&wg)
	wg.Wait()
	fmt.Println("end")
}
