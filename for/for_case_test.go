package _for

import (
	"fmt"
	"testing"
	"time"
)

//问题：for select时，如果通道已经关闭会怎么样？如果select中只有一个case呢？
//结果：当未有值时，由于缓存为0，<-ch 没数据，则一直走default，当有值输入后，但是在通道关闭后，这个通道一直能读出内容。
func TestForCase(t *testing.T) {
	ch := make(chan int)

	go func() {
		time.Sleep(time.Second * 1)
		ch <- 1
		close(ch)
	}()
	for {
		select {
		case a, ok := <-ch:
			fmt.Printf("fmt:%d time:%v  %v \t\n", a, time.Now(), ok)
			time.Sleep(500 * time.Millisecond)
		//default:
		//	fmt.Println("咩有读出来")
		//	time.Sleep(500 * time.Millisecond)
		}

	}

}

//问题：怎么样才能不读关闭后通道
func TestForCase2(t *testing.T) {
	ch := make(chan int)

	go func() {
		time.Sleep(time.Second * 1)
		ch <- 1
		close(ch)
	}()
	for {
		select {
		case a, ok := <-ch:
			fmt.Printf("fmt:%d time:%v  %v \t\n", a, time.Now(), ok)
			time.Sleep(500 * time.Millisecond)
			if !ok {
				ch = nil //把关闭后的通道复值为nil，则select读取则会阻塞
			}
		//default:
		//	fmt.Println("咩有读出来")
		//	time.Sleep(500 * time.Millisecond)
		}

	}

}
