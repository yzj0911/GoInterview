package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func consumer(product chan int, name string, size int) {
	for {
		count, ok := <-product
		if count > 0 {
			if count >= size {
				count = count - size
				fmt.Printf("经销商"+name+"购买了%d辆汽车，当前存量%d\n", size, count)
			} else {
				size = count
				fmt.Printf("经销商"+name+"购买了%d辆汽车,工厂已经开始缺货\n", size)
				count = 0
			}

		}
		if ok {
			product <- count
			//进货间隔需要10秒钟
			time.Sleep(10 * time.Second)
		}

	}

}

func factory(product chan int, done chan bool) {
	yield := 0
	for {
		select {
		case count := <-product:
			if count == 0 {
				yield++
			}
			if yield >= 3 {
				fmt.Println("超级工厂产量不足,工厂开始考虑扩厂,新厂房建设中....")
				//close(product)
				done <- true
				//break
			}
			if count > 0 && count < 200 {
				count = count + 5
				fmt.Printf("超级工厂努力生产中....，当前汽车存量%d\n", count)
			} else if count <= 0 {
				//提高产能
				count = count + 10
				fmt.Println("仓库空了，正在加紧生产汽车中")
			} else {
				fmt.Printf("仓库满了，当前汽车容量%d\n", count)
			}
			product <- count
			//生产一个批次汽车需要1秒钟
			time.Sleep(1 * time.Second)

		default:
			product <- 1
			fmt.Println("工厂即将开始生产")
		}
	}
}

func main() {

	comsumerCout := 20
	product := make(chan int, comsumerCout)
	bools := make(chan bool)
	go factory(product, bools)
	//先让工厂生产一会
	time.Sleep(10 * time.Second)
	//现在经销商开始拿货
	for i := 1; i < comsumerCout; i++ {
		go consumer(product, strconv.Itoa(i), rand.Intn(10)+5)
		//每个客户分批次拿货
		time.Sleep(1 * time.Second)
	}
	for {
		if <-bools {
			time.Sleep(15 * time.Second)
			break
		}
	}
}
