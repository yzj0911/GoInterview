package main

import (
	"execlt1/pgbar"
	"sync"
	"testing"
	"time"
)

//func TestA(t *testing.T){
//	str := "问11"
//	StrFilterNonChinese(&str)
//	fmt.Println(str)
//}
////
//var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$")
//
////func main() {
////	str := "问11"
////	StrFilterNonChinese(&str)
////	fmt.Println(str)
////}
//
//func StrFilterNonChinese(src *string) {
//	strn := ""
//	for _, c := range *src {
//		if hzRegexp.MatchString(string(c)) {
//			strn += string(c)
//		}
//	}
//
//	*src = strn
//}

// Print text heart
// Author: ShixiangWang
// LICENSE: MIT
// Reference: https://blog.csdn.net/su_bao/article/details/80355001
func TestAiXin(t *testing.T) {

	pgb := pgbar.New("多线程进度条")
	pgbar.Println("进度条1")
	b := pgb.NewBar("1st", 20000)
	pgbar.Println("进度条2")
	b2 := pgb.NewBar("2st", 10000)
	pgbar.Println("进度条3")
	b3 := pgb.NewBar("3st", 30000)

	b.SetSpeedSection(900, 100)
	b2.SetSpeedSection(900, 100)
	b3.SetSpeedSection(900, 100)

	pgbar.Println("独立进度条")
	b4 := pgbar.NewBar(0, "4st", 4000)
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		for i := 0; i < 20000; i++ {
			b.Add()
			time.Sleep(time.Second / 2000)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			b2.Add()
			time.Sleep(time.Second / 1000)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 30000; i++ {
			b3.Add()
			time.Sleep(time.Second / 3000)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 4000; i++ {
			b4.Add()
			time.Sleep(time.Second / 300)
		}
	}()
	wg.Wait()
}
