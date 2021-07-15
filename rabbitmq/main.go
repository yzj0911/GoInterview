package main

import "github.com/streadway/amqp"
func main(){
	conn, err := amqp.Dial(amqp://guest:guest@172.17.84.205:5672/)
}
