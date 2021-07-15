package main

import (
	"fmt"
	"time"
)

// 返回生成自然数序列的管道: 2, 3, 4, ...
func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// 管道过滤器: 删除能被素数整除的数
func PrimeFilter(in <-chan int, prime int) chan int {

	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func worker(channel chan bool) {
	//for {
	//	select {
	//	default:
	//		fmt.Println("zc")
	//	case <-channel:
	//		fmt.Println("return")
	//		return
	//	}
	//}
	for {
		select {
		default:
			//channel 未退出
			fmt.Println("2")
		case <-channel:
			//接受到channel 表示退出
			fmt.Println("end")
			return
		}
	}
}

func main() {
	cha := make(chan bool, 1)
	//defer close(cha)
	go func() {
		for i := 0; i <= 10; i++ {
			if i==3{
				cha <- true
			}
			time.Sleep(time.Second * 1)
			fmt.Println("1111111")
		}
	}()
	go func() {
		fmt.Println("12312")
		for i := 0; i <= 100; i++ {
			//if i == 100 {
			//	cha <- true
			//}
			time.Sleep(time.Second * 1)
			fmt.Println("22")
		}
	}()
	fmt.Println("======")
	a := <-cha

	fmt.Println("-----")
	//a := <-cha
	fmt.Println(a)
	//go worker(cha)
	//time.Sleep(time.Second * 3)
	//cha <- true
}
