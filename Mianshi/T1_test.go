package Mianshi

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
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
		m[i] = &i
	}

	for i, _ := range m {
		print(i)
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
const N3=10


func Test5(t *testing.T) {
	m:=make(map[int]int)
	wg:=&sync.WaitGroup{}
	wg.Add(N3)
	for i:=0;i<N3;i++{
		go func() {
			defer wg.Done()
			m[rand.Int()]=rand.Int()
		}()
	}
	wg.Wait()
	fmt.Println(len(m))
}

//-----------------------------------------------------

type S3 struct{
	a,b,c string
}

func Test6(t *testing.T) {
	x:=interface{}(S3{a:"1",b:"2",c:"3"})
	y:=interface{}(S3{a:"1",b:"2",c:"3"})
	fmt.Println(x==y)
}

//------------------------------------------------------
type S7 struct {
	Name string
}
func Test7(t *testing.T) {
	m:=map[string]*S7{"x":&S7{Name:"one"}}

	//m["x"]=S7{Name:"two"}
	m["x"].Name="two"
}

//------------------------------------------------

type Result struct {
	//Status int
	status int
}
func Test8(t *testing.T) {
	var data =[]byte(`{"status":200}`)
	resutl:=&Result{

	}
	if err:=json.Unmarshal(data,resutl);err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("restult=%+v\n",resutl)
}


