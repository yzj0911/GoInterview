package Mianshi

import (
	"context"
	reand2 "crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"reflect"
	"runtime"
	"sync"
	"testing"
)

func TestMai(t *testing.T) {
	go func() {
		runtime.GC()
	}()
	_ = new(int64)
	_ = make([]int64, 10)
	_ = context.Background()
	//fmt.Println(GoID())
}

///////////=========================================================================================================
const N = 3

func Test1(t *testing.T) {
	m := make(map[int]*int)
	for i := 0; i < N; i++ {
		var s = i
		m[i] = &s
	}

	for _, i := range m {
		print(*i)
	}
}

//-------------------------------------------------

type S struct {
	m string
}

func f() *S {
	return &S{"foo"}
}
func Test2(t *testing.T) {
	p := f()
	println(p.m)
}

//--------------------------------------------------

const N2 = 10

func Test3(t *testing.T) {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N2)
	for i := 0; i < N2; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	println(len(m))
}

//--------------------------------------------------

type S1 struct{}

func (S S1) f() {
	fmt.Println("S1.f()")
}

func (S S1) g() {
	fmt.Println("S1.g()")
}

type S2 struct {
	S1
}

func (S S2) f() {
	fmt.Println("S2.f()")
}

type I interface {
	f()
}

func printType(i I) {
	if s1, ok := i.(S1); ok {
		s1.f()
		s1.g()
	}
	if s1, ok := i.(S2); ok {
		s1.f()
		s1.g()
	}
}
func Test4(t *testing.T) {
	printType(S1{})
	printType(S2{})
}

//---------------------------------------------
const N3 = 10

func Test5(t *testing.T) {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	wg.Add(N3)
	for i := 0; i < N3; i++ {
		go func() {
			defer wg.Done()
			m[rand.Int()] = rand.Int()
		}()
	}
	wg.Wait()
	fmt.Println(len(m))
}

//-----------------------------------------------------

type S3 struct {
	a, b, c string
}

func Test6(t *testing.T) {
	x := interface{}(S3{a: "1", b: "2", c: "3"})
	y := interface{}(S3{a: "1", b: "2", c: "3"})
	fmt.Println(x == y)
}

//------------------------------------------------------
type S7 struct {
	Name string
}

func Test7(t *testing.T) {
	m := map[string]*S7{"x": &S7{Name: "one"}}

	//m["x"]=S7{Name:"two"}
	m["x"].Name = "two"
}

//------------------------------------------------

type Result struct {
	//Status int
	status int
}

func Test8(t *testing.T) {
	var data = []byte(`{"status":200}`)
	resutl := &Result{}
	if err := json.Unmarshal(data, resutl); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("restult=%+v\n", resutl)
}

//-------------------------------------
type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {
	fmt.Println("student")
}

func live() People {
	var stu *Student
	return stu
}

func Test9(t *testing.T) {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		live().Show()
		//fmt.Println()
		fmt.Println("BBBBBBB")
	}
}

//=========================================================
type A interface {
	Say(string) (string, int)
}

type Test10One struct {
	Name string
	Age  int
}

func (one *Test10One) Say(food string) (string, int) {
	fmt.Println(food)
	return one.Name, one.Age
}

func Test10(t *testing.T) {
	t10 := &Test10One{
		Name: "yyy",
		Age:  10,
	}
	t10T := reflect.TypeOf(t10)
	fmt.Println(t10T.Elem().NumIn())
	method, ok := t10T.MethodByName("Say")
	if !ok {
		fmt.Println("not ok")
		return
	}

	a := method.Func.Call([]reflect.Value{reflect.ValueOf(t10), reflect.ValueOf("shot")})
	fmt.Println(a)
}

func Test11(t *testing.T) {
	for i := 0; i < 20; i++ {
		fmt.Print(Y11(), ",")
	}
}
func x11() int {
	b, _ := reand2.Int(reand2.Reader, big.NewInt(10_001))
	if b.Int64() >= 50_00 {
		return 1
	}
	return 0
}

func Y11() int {
	ret := 0
	for i := 0; i < 4; i++ {
		ret |= x11() << i
	}
	if ret > 9 {
		ret = Y11()
	}
	return ret
}
