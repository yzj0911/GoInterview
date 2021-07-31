package main

import (
	"fmt"
	"reflect"
)

func main() {
	s1 := struct {
		id int
		m  map[string]string
	}{id: 1, m: map[string]string{"a": "1"}}
	s2 := struct {
		id int
		m  map[string]string
	}{id: 1, m: map[string]string{"a": "1"}}

	//if s1 == s2 {
	//	fmt.Println("great")
	//} 编译失败
	if reflect.DeepEqual(s1, s2) {
		fmt.Println("great")
	}
	fmt.Println("end")
}
