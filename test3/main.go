package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func main() {

	a := SHA1Pre8Hex("你好")
	fmt.Println(a)
	b := SHA1Pre8Hex("你好")
	fmt.Println(b)

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
