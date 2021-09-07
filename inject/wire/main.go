package main

import "fmt"

type apple struct {
	name  string
	score int
}

func (a *apple) say() string {
	return a.name
}

type banana struct {
	job   string
	price int
}

func (b *banana) speak() string {
	return b.job
}

type shop struct {
	a *apple
	b *banana
}

func (s *shop) sail() string {
	return s.a.say() + s.b.speak()
}

func NewA() *apple {
	return &apple{
		name:  "apple",
		score: 1,
	}
}

func NewB() *banana {
	return &banana{
		job:   "banana",
		price: 2,
	}
}

func NewS(a *apple, b *banana) shop {
	return shop{a: a, b: b}
}

func Init1() {
	a := NewA()
	b := NewB()
	s := NewS(a, b)
	fmt.Println(s.sail())
}

func main() {
	//Init1()
	s := InitializeShop()
	sail := s.sail()
	fmt.Println(sail)
}
