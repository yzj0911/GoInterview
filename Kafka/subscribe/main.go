package main

import (
	"github.com/zngw/kafka"
	"github.com/zngw/log"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	err := log.Init("", nil)
	if err != nil {
		panic(err)
	}
	err = kafka.InitConsumer("127.0.0.1:9092")
	if err != nil {
		panic(err)
	}
	go func() {
		err = kafka.LoopConsumer("Test", TopicCallBack)
		if err != nil {
			panic(err)
		}
	}()
	signal.Ignore(syscall.SIGHUP)
	runtime.Goexit()
}

func TopicCallBack(data []byte) {
	log.Trace("kafka", "Test:"+string(data))
}
