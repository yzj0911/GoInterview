// discovery.go
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	prefix = "register"
	client *clientv3.Client
)

var (
	port     = flag.Int("port", 30001, "service port")
	endpoint = flag.String("endpoints", "http://127.0.0.1:2379", "etcd endpoints")
)

type SvConfig struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

func watcher() error {
	var err error
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(*endpoint, ","),
		DialTimeout: time.Second * 3,
	})
	if err != nil {
		return fmt.Errorf("connect etcd cluster failed: %v", err.Error())
	}

	go func() {
		resp := client.Watch(context.TODO(), prefix, clientv3.WithPrefix())
		for ch := range resp {
			for _, event := range ch.Events {
				switch event.Type {
				case clientv3.EventTypePut:
					if event.IsCreate() {
						srv := parseSrv(event.Kv.Value)
						log.Printf("discovery service %s at %s:%d", srv.Name, srv.Host, srv.Port)
					}
				case clientv3.EventTypeDelete:
					log.Printf("delete service %s", event.Kv.Key)
				}
			}
		}
	}()

	return err
}

func parseSrv(text []byte) *SvConfig {
	svc := &SvConfig{}
	json.Unmarshal(text, &svc)
	return svc
}

func main() {
	flag.Parse()

	// 绑定服务地址和端口
	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
	if err != nil {
		panic(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		<-ch
		os.Exit(1)
	}()

	watcher()

	log.Printf("discovery start at %d", *port)
	// server todo
	for {
		lis.Accept()
	}
}