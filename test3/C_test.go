package main

import (
	"fmt"
	"testing"
)

type c struct {
}

func TestCA(t *testing.T) {
	//vipRes, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(18)/float64(31)), 64)
	//fmt.Println(math.Round(vipRes) * 100)
	//
	//a := make([]int64, 1, 10)
	//a[0] = 1
	//a = append(a, 12)
	//
	//var d = &c{}
	//fmt.Println(d)
	//
	//fmt.Println(a)
	//
	//m := make(map[string]int, 10)
	//m["one"] = 1
	//fmt.Println(m)
	//var a = []int64{1, 2, 3}
	//for i, v := range a {
	//	fmt.Println(i)
	//	if i == len(a) {
	//		panic("111")
	//	}
	//	fmt.Println(v)
	//}
	a := make([]int64, 0, 6)

	a = append(a, 10)
	fmt.Println(a)
}
