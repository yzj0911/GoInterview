package main

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	ch := GenerateNatural()
	for i := 0; i <= 100; i++ {
		pri := <-ch
		fmt.Println(i, "  ", pri)
		ch = PrimeFilter(ch, pri)
	}
}

func TestSelect(t *testing.T) {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0:
				fmt.Println("0")
			case ch <- 1:
				fmt.Println("1")
			}
		}
	}()
	for v := range ch {
		fmt.Println(v)
	}
}
func workers(wg *sync.WaitGroup, channel chan bool) {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("====")
		case <-channel:
			fmt.Println("return ")
			return
		}
	}
}

func TestChannel(t *testing.T) {
	ch := make(chan bool)
	//sync.waitGroup 相当于计数器
	var wg sync.WaitGroup
	for i := 0; i <= 100; i++ {
		//添加一个wg执行
		wg.Add(1)
		go workers(&wg, ch)
	}
	time.Sleep(1 * time.Second)
	close(ch)
	//等待所有都wg 执行完毕
	wg.Wait()
}

func workContext(ctx context.Context, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	//定时器
	timer := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timer.C:
			fmt.Println("----")
		default:
			fmt.Println("work", i)
		case <-ctx.Done():
			fmt.Println("return", i)
			return
		}
	}
}

func TestCannelContext(t *testing.T) {
	//ctx, channel := context.WithTimeout(context.Background(), 10*time.Second)
	//var wg sync.WaitGroup
	//for i := 0; i <= 100; i++ {
	//	wg.Add(1)
	//	go workContext(ctx, &wg)
	//}
	//time.Sleep(time.Second * 10)
	//channel()
	//
	//wg.Wait()

	ctx, channel := context.WithTimeout(context.Background(), 10*time.Second)
	var wg sync.WaitGroup
	for i := 0; i <= 4; i++ {
		wg.Add(1)
		time.Sleep(1 * time.Second)
		go workContext(ctx, &wg, i)
	}
	fmt.Println("end ======")
	channel()
	wg.Wait()
}

func TestChannelNull(t *testing.T) {
	cha := make(chan int, 0)
	cha <- 1
	fmt.Println("end")
}
