package main

import (
	"github.com/Shopify/sarama"
	"github.com/zngw/log"
	"strings"
)

var consumer sarama.Consumer

// ConsumerCallback 消费者回调用
type ConsumerCallback func(data []byte)

func initConsumer(hosts string) {
	config := sarama.NewConfig()
	client, err := sarama.NewClient(strings.Split(hosts, ""), config)
	if err != nil {
		log.Error(err)
	}
	consumer, err = sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Error(err)
	}
}

// LoopConsumer 消费者循环
func LoopConsumer(topic string, callBack ConsumerCallback) {
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Error(err)
		return
	}
	defer partitionConsumer.Close()
	for {
		mes := <-partitionConsumer.Messages()
		if callBack != nil {
			callBack(mes.Value)
		}
	}
}

func Close() {
	if consumer != nil {
		consumer.Close()
	}
}
