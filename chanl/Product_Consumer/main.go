package main

import (
	"fmt"
)

//生产者消费者
//生产者：生成 多个序列
func Product(p int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * p
	}
}

//消费者 读取in队列中的所有 int
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 64)
	//time.ParseDuration()
	go Product(1, ch)
	go Product(1, ch)
	go Consumer(ch)
	//signal包 一个是notify方法用来监听收到的信号；一个是 stop方法用来取消监听。
	//signal.Stop(c) //不允许继续往c中存入内容
	//s := <-c
	//sig := make(chan os.Signal, 1)
	//signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	//signal.Stop(sig)
	//a := <-sig
	//fmt.Println(a)
	//fmt.Printf("quit (%v)\n", <-sig)
	//time.Sleep(time.Second * 5)
}
