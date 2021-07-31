package main

import (
	"fmt"
	"reflect"
)

func Test(is ...interface{}) {
	for _, i := range is {
		if reflect.TypeOf(i).String() != "string" {
			fmt.Println(i)
		}

	}
}

func main() {
	var is []interface{}
	Test(is)
	i := GetVale()
	switch i.(type) {
	case int:
		fmt.Println(i)
	default:
		fmt.Println("Err")
	}
}

func GetVale() interface{} {
	return 1
}
