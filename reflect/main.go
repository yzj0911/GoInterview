package main

import (
	"fmt"
	"reflect"
)

type Animal interface {
	Speak()
}

type Cat struct {
	name string
}

func (c Cat) Speak() {
	fmt.Println("MM")
}

type Dog struct {
	age int32
}

func (d Dog) Speak() {
	fmt.Println(d.age, "  jj")
}

func main() {
	var animal Animal
	animal = Cat{}
	animal.Speak()
	//var c interface{}
	s, ok := animal.(Cat)
	if !ok {
		fmt.Println(ok)
		return
	}
	s.Speak()
	fmt.Println("end")

	fmt.Println(reflect.TypeOf(s).String())
	//params := make([]reflect.Value,0)
	method := reflect.ValueOf(s).MethodByName("Speak")

	method.Call(nil)

	var a interface{}
	i,ok:=a.(Dog)
	if !ok{
		fmt.Println("not ok")
		return
	}
	reflect.ValueOf(i).MethodByName("Speak").Call(nil)


}
