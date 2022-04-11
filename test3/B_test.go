package main

import (
	"fmt"
	"sort"
	"testing"
)

func t(c <-chan string) (string, bool) {
	a, ok := <-c
	return a, ok
}

func Test_T(t2 *testing.T) {
	a := 1
	defer fmt.Println("the value of a1:", a)
	a++

	defer func() {
		fmt.Println("the value of a2:", a)
	}()
}

func Test_T2(t *testing.T) {
	a := make([]int, 4, 6)
	fmt.Println(a)
	b := a
	b = append(b, 1)
	fmt.Println(b)
	a[1] = 1
	fmt.Println(b)
}

func Test_T3(t *testing.T) {
	a := []int{2, 3, 1, 0, 4, 5, 6, 7}
	b := a[2:5:8]
	fmt.Println(b)
	fmt.Println(cap(b), len(b))
	sort.Ints(a)
	fmt.Println(a)
	fmt.Println(b)
}
