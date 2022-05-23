package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://123:123@localhost:5672")
	fmt.Println(err)
	defer conn.Close()

	ch, err := conn.Channel()
	fmt.Println(err)
	defer ch.Close()

	args := amqp.Table{"x-dead-letter-exchange": "dlx"}
	q, err := ch.QueueDeclare("test", true, false, false, false, args) // 声明一个test队列，并设置队列的死信交换机为"dlx"

	body := "hello world1"
	for i := 0; i < 10; i++ {
		err = ch.Publish("", q.Name, false, false, amqp.Publishing{
			Body:       []byte(body),
			Expiration: "5000", // 设置TTL为5秒
		})
		fmt.Println(err)
	}
}
