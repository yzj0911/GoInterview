package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func main() {
	//a := fmt.Sprintf("0 0 1/%d * ?", 1)
	//c, e := cron.ParseStandard(a)
	//if e != nil {
	//	fmt.Println(e)
	//}
	//nextTime := c.Next(time.Now())
	//
	//b := fmt.Sprintf("0 0 1/%d * ?", 2)
	//cs, err := cron.ParseStandard(b)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//nextTimes := cs.Next(time.Now())
	//fmt.Println(nextTime)
	//fmt.Println(nextTimes)
	con:=fmt.Sprintf("@every %dh",24*5)
	c, e := cron.ParseStandard(con)
	if e != nil {
		fmt.Println(e)
	}
fmt.Println(c)
	nextTime := c.Next(time.Now())
	fmt.Println(nextTime)
}
