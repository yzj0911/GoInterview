package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/smallnest/rpcx/log"
	"time"
	"unsafe"
)

type SubscribeCallBack func(channel, message string)

type Subscriber struct {
	clint redis.PubSubConn
	//存放回调函数
	cbMap map[string]SubscribeCallBack
}

func (c *Subscriber) Connect(ip string, port uint16) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
	}
	c.clint = redis.PubSubConn{conn}
	c.cbMap = make(map[string]SubscribeCallBack)
	go func() {
		for {
			fmt.Println("=====wait====")
			switch res := c.clint.Receive().(type) {
			case redis.Message:
				channel := (*string)(unsafe.Pointer(&res.Channel))
				message := (*string)(unsafe.Pointer(&res.Data))
				c.cbMap[*channel](*channel, *message)
			case redis.Subscription:
				fmt.Printf("%s %s :%d \n", res.Channel, res.Kind, res.Count)
			case error:
				fmt.Println("error========")
				continue
			}
		}
	}()
}

func (c Subscriber) Close() {
	err := c.clint.Close()
	if err != nil {
		fmt.Println(err, "\n")
	}
}

func (c Subscriber) Subscribe(channel interface{}, cb SubscribeCallBack) {
	err := c.clint.Subscribe(channel)
	if err != nil {
		fmt.Println(err)
	}
	c.cbMap[channel.(string)] = cb
}

func TestCallback1(chann, meg string) {
	log.Debug("TestCallback1 channel : ", chann, " Message ", meg)
}

func TestCallback2(chann, meg string) {
	log.Debug("TestCallback2 channel : ", chann, " Message ", meg)
}

func TestCallback3(chann, meg string) {
	log.Debug("TestCallback3 channel : ", chann, " Message ", meg)
}

func main() {
	log.Infof("\n===============main start ===============")
	var sub Subscriber
	sub.Connect("loclhost", 6397)
	sub.Subscribe("test_chan1", TestCallback1)
	sub.Subscribe("test_chan2", TestCallback2)
	sub.Subscribe("test_chan3", TestCallback3)
	for {
		time.Sleep(3 * time.Minute)
	}

}
