package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"reflect"
)

const (
	a1 int = iota
	a2
	a3
)

func main() {

	aa := SHA1Pre8Hex("你好")
	fmt.Println(aa)
	bb := SHA1Pre8Hex("你好")
	fmt.Println(bb)

	a := []int{1, 2, 3, 4}
	b := []int{1, 3, 2, 4}
	c := []int{1, 2, 3, 4}
	fmt.Println(reflect.DeepEqual(a, b))
	fmt.Println(reflect.DeepEqual(a, c))
	fmt.Printf("%v\n", a)
	fmt.Printf("%+v\n", a)
}

// SHA1Pre8Hex 获取字符串经过SHA1，再转化为Hex后的前8位
func SHA1Pre8Hex(data string) string {
	var tmp = sha1.Sum([]byte(data))
	return hex.EncodeToString(tmp[:4])
}

func BinarySearch(list []int, item int) int {
	var low int
	var hight int = len(list) - 1
	for ; low <= hight; {
		mid := (low + hight) / 2
		guss := list[mid]
		if guss < item {
			low = mid + 1
		} else if guss == item {
			return mid
		} else {
			hight = mid - 1
		}

	}
	return low
}
