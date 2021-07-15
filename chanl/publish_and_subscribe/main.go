package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

)

//订阅发布
type (
	subscriber chan interface{}         //订阅者为一个管道
	topicFunc  func(v interface{}) bool //主题为一个过滤器
)

// Publish 发布者
type Publish struct {
	m           sync.RWMutex             //读写锁
	buffer      int                      //订阅队列的缓存大小
	timeout     time.Duration            //发布超时时间
	subscribers map[subscriber]topicFunc //订阅者信息
}

// NewPublisher 新建发布者
func NewPublisher(publishTimeout time.Duration, buffer int) *Publish {
	return &Publish{
		buffer:      buffer,
		timeout:     publishTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// Subscribe 添加一个新的订阅者，订阅全部主题
func (p *Publish) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

// SubscribeTopic 添加一个新的订阅者，订阅过滤器筛选后的主题
func (p *Publish) SubscribeTopic(topicFunc topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topicFunc
	p.m.Unlock()
	return ch
}
// 退出订阅
func (p *Publish) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, sub)
	close(sub)
}
// 发布一个主题
func (p *Publish) Publish(v interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

// 关闭发布者对象，同时关闭所有的订阅者管道。
func (p *Publish) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

// 发送主题，可以容忍一定的超时
func (p *Publish) sendTopic(
	sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup,
) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}

	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

/*
在发布订阅模型中，每条消息都会传送给多个订阅者。
发布者通常不会知道、也不关心哪一个订阅者正在接收主题消息。
订阅者和发布者可以在运行时动态添加，是一种松散的耦合关系，这使得系统的复杂性可以随时间的推移而增长。
在现实生活中，像天气预报之类的应用就可以应用这个并发模式。
 */
func main() {
	p := NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})
	p.Publish("hello,  world!")
	p.Publish("hello, golang!")

	go func() {
		for  msg := range all {
			fmt.Println("all:", msg)
		}
	} ()

	go func() {
		for  msg := range golang {
			fmt.Println("golang:", msg)
		}
	} ()

	// 运行一定时间后退出
	time.Sleep(3 * time.Second)
}