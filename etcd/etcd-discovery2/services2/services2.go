package main
import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

const (
	EtcdPrefix   = "/test/server/"
	ServerSerial = "2"
	Address      = "http://127.0.0.1:18082/"
)

var (
	EtcdAddress = []string{"http://47.99.94.72:2379"}
	leaseTTL    = 5
)

type HealthProvider struct {
	etcdClient *EtcdClient
}

var (
	healthProvider     *HealthProvider
	healthProviderOnce sync.Once
)

func GetHealthProvider() *HealthProvider {
	healthProviderOnce.Do(func() {
		healthProvider = &HealthProvider{
			etcdClient: NewEtcdClient(),
		}
	})
	return healthProvider
}

type EtcdClient struct {
	address  []string
	username string
	password string
	kv       clientv3.KV
	client   *clientv3.Client
	ctx      context.Context
	lease    clientv3.Lease
	leaseID  clientv3.LeaseID
	leaseTTL int64
}

func NewEtcdClient() *EtcdClient {
	var client = &EtcdClient{
		ctx:      context.Background(),
		address:  EtcdAddress,
		leaseTTL: int64(leaseTTL),
	}
	err := client.connect()
	if err != nil {
		panic(err)
	}
	return client
}
//连接到etcd
func (etcdClient *EtcdClient) connect() (err error) {
	etcdClient.client, err = clientv3.New(clientv3.Config{
		Endpoints:   etcdClient.address,
		DialTimeout: 5 * time.Second,
		TLS:         nil,
		Username:    etcdClient.username,
		Password:    etcdClient.password,
	})
	if err != nil {
		return
	}
	etcdClient.kv = clientv3.NewKV(etcdClient.client)
	etcdClient.ctx = context.Background()
	return
}

func (etcdClient *EtcdClient) Close() (err error) {
	return etcdClient.client.Close()
}

//注册，写入到etcd中
func (etcdClient *EtcdClient) register(address string) (*clientv3.PutResponse, error) {
	etcdClient.lease = clientv3.NewLease(etcdClient.client)
	leaseResp, err := etcdClient.lease.Grant(etcdClient.ctx, etcdClient.leaseTTL)
	if err != nil {
		return nil, err
	}
	etcdClient.leaseID = leaseResp.ID
	return etcdClient.kv.Put(etcdClient.ctx, EtcdPrefix+ServerSerial, address, clientv3.WithLease(leaseResp.ID))
}

//LeaseKeepAlive 定时检测生命周期并重新赋值
func (etcdClient *EtcdClient) LeaseKeepAlive() error {
	if etcdClient.lease == nil {
		_, err := etcdClient.register(Address)
		if err != nil {
			return err
		}
	}
	_, err := etcdClient.lease.KeepAlive(etcdClient.ctx, etcdClient.leaseID)
	if err != nil {
		return err
	}
	return nil
}

func healthCheck(provider *HealthProvider) {
	var tick = time.NewTicker(time.Second)
	for {
		select {
		case <-tick.C:
			err := provider.etcdClient.LeaseKeepAlive()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}
}

func main() {

	provider := GetHealthProvider()
	go healthCheck(provider)

	defer provider.etcdClient.Close()

	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "two")
	})

	engine.Run(":18082")
}