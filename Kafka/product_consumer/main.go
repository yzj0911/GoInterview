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

	err = kafka.InitProducer("127.0.0.1:9092")
	if err != nil {
		panic(err)
	}
	defer kafka.Close()
	kafka.Send("Test", "This is Test Msg")
	kafka.Send("Test", "Hello Guoke")
	signal.Ignore(syscall.SIGHUP)
	runtime.Goexit()
}
