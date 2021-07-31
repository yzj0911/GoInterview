package main

import (
	"context"
	"fmt"
	"sync"
	time2 "time"
)

func main() {
	ct, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		times := time2.NewTimer(time2.Second * 2)
		defer times.Stop()
		var i = 1
		select {
		case <-ct.Done():
			fmt.Println("s")
			return
		case <-times.C:
			i++
		default:
			if i >= 19 {
				cancel()
			}
		}
	}()
}

type ExcusiveQueue struct {
	queued     int32
	singleChan chan struct{}
}

type ExclusiveMap struct {
	sync.Mutex
	resourceMap map[interface{}]*ExcusiveQueue
}

//获得一个锁通道
func (e *ExclusiveMap) acquire(key interface{}) chan struct{} {
	e.Lock()
	defer e.Unlock()
	value, b := e.resourceMap[key]
	if b {
		value.queued++
	} else {
		value = &ExcusiveQueue{queued: 1, singleChan: make(chan struct{}, 1)}
		e.resourceMap[key] = value
	}
	return value.singleChan
}

//释放排他锁
func (e *ExclusiveMap) release(key interface{}) {
	e.Lock()
	defer e.Unlock()
	value, b := e.resourceMap[key]
	if b {
		value.queued--
		if value.queued == 0 {
			delete(e.resourceMap, key)
		}
		<-value.singleChan
	}
}
