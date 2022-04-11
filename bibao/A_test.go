package main

import (
	"fmt"
	"testing"
)

//面试

func Test(t *testing.T) {
	var funcSlice []func()
	for i := 0; i < 3; i++ {
		//println(&i) // 0xc0000ac1d0 0xc0000ac1d0 0xc0000ac1d0
		funcSlice = append(funcSlice, func() {
			//println(&i)
			println(i)
		})

	}
	for j := 0; j < 3; j++ {
		funcSlice[j]() // 3, 3, 3
	}
}

type fn func(x int) (y int)

func func1() []fn {
	var funlist []fn
	for i := 0; i < 4; i++ {
		var func2 = func(x int) (y int) {
			return i * x
		}
		funlist = append(funlist, func2)
	}
	return funlist
}
func Test_ff(t *testing.T) {
	for _, fn := range func1() {
		fmt.Println(fn(1))
		fmt.Println(fn(2))
	}
}