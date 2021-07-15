package main

import "fmt"

//map 定义了需要初始化
type A struct {
	a string
	b string
	c int
	d map[int]B //	a.d = make(map[int]B)
}

type B struct {
}

func a() {
}

func main() {
	var a A
	var b B
	a.d = make(map[int]B)
	a.d[1] = b
	fmt.Println(a)
}
