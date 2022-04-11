package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"log"
	"time"
)

var (
	Prefix      = "/etcd"
	EtcdAddress = "http://47.99.94.72:2379"
	Client      *clientv3.Client
)

//生产者
func publisher(client *clientv3.Client) {
	go func() {
		timer := time.NewTimer(time.Second)
		for range timer.C {
			now := time.Now()
			key := fmt.Sprintf("%s/%d", Prefix, now.Second())
			value := now.String()
			client.Put(context.TODO(), key, value)
		}
	}()

}

//消费者
func subscriber(client *clientv3.Client) {
	watcher := client.Watch(context.TODO(), Prefix, clientv3.WithPrefix())
	for ch := range watcher {
		for _, e := range ch.Events {
			if e.IsCreate() {
				log.Printf("received %s => %s\n", e.Kv.Key, e.Kv.Value)
			}
		}
	}
}

func main() {
	flag.Parse()
	var err error
	Client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{EtcdAddress},
		DialTimeout: time.Second * 2,
	})
	if err != nil {
		log.Fatalln("connect etcd cluster failed: " + err.Error())
	}
	publisher(Client)
	subscriber(Client)

	select {}
}
