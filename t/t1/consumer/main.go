package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://123:123@localhost:5672")
	fmt.Println(err)

	c, err := conn.Channel()
	fmt.Println(err)

	msgs, err := c.Consume("dlxQueue", "", true, false, false, false, nil) //监听dlxQueue队列
	fmt.Println(err)

	for d := range msgs {
		fmt.Printf("信息: %s\n", d.Body)
	}
}
