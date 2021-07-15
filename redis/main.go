package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var Pool *redis.Pool

func init() {
	//Pool = &redis.Pool{
	//	MaxIdle:     16,
	//	MaxActive:   0,
	//	IdleTimeout: 300,
	//	Dial: func() (redis.Conn, error) {
	//		return redis.Dial("tcp", "127.0.0.1:6379")
	//	},
	//}

}

func main() {
	//c := Pool.Get()
	//defer c.Close()
	//_, err := c.Do("set", "abc", 100)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//r, err := redis.Int(c.Do("get", "abc"))
	//fmt.Println(r)
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("err")
		return
	}
	defer c.Close()
	//hset
	fmt.Println("=================== HSet ===================")
	if _, err := c.Do("HSet", "book", "abc", 100); err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.Int(c.Do("HGet", "book", "abc"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
	//list
	fmt.Println("=================== List ===================")
	if _, err := c.Do("lpush", "book_list", "11", "22", "44"); err != nil {
		fmt.Println(err)
		return
	}
	listRes, err := redis.String(c.Do("lpop", "book_list"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(listRes)
	//set
	fmt.Println("=================== set ====================")
	if _, err := c.Do("set", "a", 1001111111); err != nil {
		fmt.Println(err)
		return
	}
	setRes, err := redis.Int(c.Do("get", "a"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(setRes)
	//batch_set
	fmt.Println("============= batch_Set =============")
	if _, err := c.Do("mset", "abc", 1, "a", 1); err != nil {
		fmt.Println(err)
		return
	}
	sets, err := redis.Ints(c.Do("mget", "abc", "a"))
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, set := range sets {
		fmt.Println(set)
	}
	//过期设置
	fmt.Println("=================过期设置==========")
	c.Do("expire","abc",10)
}
