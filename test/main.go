package main

// 导入系统包
import (
	"flag"
	"fmt"
)

// 定义命令行参数
var mode = flag.String("mode", "", "process mode")

func main1() {
	//var a int = 8000
	//l, err := net.ListenTCP("tcp4", &net.TCPAddr{IP: net.ParseIP("0.0.0.0"), Port: a})
	//if err != nil {
	//	return
	//}
	//fmt.Println(time.Now())
	//fmt.Println("end")
	//l.Close()
	pase_student()
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []*student{
		&student{Name: "zhou", Age: 24},
		&student{Name: "li", Age: 23},
		&student{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = stu
		fmt.Println(&stu)
	}
	fmt.Println(m)
}

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println(live().Show)
		fmt.Println("BBBBBBB")
	}
}