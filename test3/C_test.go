package main

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

type c struct {
}

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// 得到id字符串
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
func TestCA(t *testing.T) {
	go func() {
		runtime.GC()
	}()
	_ = new(int64)
	_ = make([]int64, 10)
	_ = context.Background()
	fmt.Println(GoID())
	
	//a := []string{"1A", "2B", "7C", "12D"}
	//for _, v := range a {
	//
	//	fmt.Println(v + " : " + strconv.Itoa(len(v)))
	//	for _, s := range v {
	//		fmt.Println(s)
	//	}
	//
	//}
	//sort.SliceStable(a, func(i, j int) bool { return (a[i]) > (a[j]) })
	//fmt.Println(1 << 31)
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

}
