package main

import (
	"github.com/Shopify/sarama"
	"github.com/opentracing/opentracing-go/log"
	"strings"
)

var product sarama.AsyncProducer

func InitProducer(hosts string) {
	config := sarama.NewConfig()
	client, err := sarama.NewClient(strings.Split(hosts, " "), config)
	if err != nil {
		log.Error(err)
	}
	product, err = sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		log.Error(err)
	}
}

func send(topic, data string) {
	product.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(data)}
}

func Close() {
	if product != nil {
		product.Close()
	}
}
