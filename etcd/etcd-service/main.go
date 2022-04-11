// server.go
package main

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	prefix     = "register"
	client     *clientv3.Client
	stopSignal = make(chan struct{}, 1)
	srvKey     string
)

var (
	serv     = flag.String("name", "hello", "service name")
	port     = flag.Int("port", 30000, "service port")
	endpoint = flag.String("endpoints", "http://47.99.94.72:2379", "etcd endpoints")
)

type SvConfig struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

func Register(endpoints string, config *SvConfig, interval time.Duration, ttl int) error {
	// 解析服务端的值
	srvValue, _ := json.Marshal(config)
	srvKey = fmt.Sprintf("%s/%x", prefix, md5.Sum(srvValue))
fmt.Println(srvKey)
	var err error
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(endpoints, ","),
		DialTimeout: time.Second * 2,
	})
	if err != nil {
		return fmt.Errorf("register service failed: %v", err)
	}

	go func() {
		timer := time.NewTicker(interval)
		for {

			resp, _ := client.Grant(context.TODO(), int64(ttl))

			_, err = client.Get(context.TODO(), srvKey)
			if err != nil {
				// 捕获 key 不存在的场合
				if errors.Is(err, rpctypes.ErrKeyNotFound) {
					_, err = client.Put(context.TODO(), srvKey, string(srvValue), clientv3.WithLease(resp.ID))
					if err != nil {
						log.Printf("register service %s at %s:%d\n", config.Name, config.Host, config.Port)
					}
				}
			} else {
				// 如果key存在就更新ttl
				_, err = client.Put(context.TODO(), srvKey, string(srvValue), clientv3.WithLease(resp.ID))
			}
			select {
			case <-stopSignal:
				return
			case <-timer.C:
			}
		}
	}()

	return err
}

func Unregister() error {
	stopSignal <- struct{}{}
	stopSignal = make(chan struct{}, 1)
	_, err := client.Delete(context.TODO(), srvKey)
	return err
}

func main() {
	flag.Parse()

	// 绑定服务地址和端口
	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
	if err != nil {
		panic(err)
	}

	config := &SvConfig{
		Name: *serv,
		Host: "127.0.0.1",
		Port: *port,
	}
	Register(*endpoint, config, time.Second*3, 15)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		<-ch
		Unregister()
		os.Exit(1)
	}()

	log.Printf("service %s start at %d", *serv, *port)
	// server todo
	for {
		lis.Accept()
	}
}
