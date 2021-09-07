package Test

import (
	"fmt"
	"testing"
)

type People interface {
	Show()
}

type Student struct {}

func (stu *Student) Show() {
}
func live() People {
	var stu *Student
	return stu
}

func Test_WithCancel(t *testing.T) {
	total := 0
	var sum int
	for i := 1; i <= 10; i++ {
		sum += i
		go func() {
			total += i
		}()
	}
	fmt.Printf("total:%d sum %d \n", total, sum)
	fmt.Println(live())
}

