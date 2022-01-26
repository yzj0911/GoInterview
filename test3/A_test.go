package main

import (
	"fmt"
	"testing"
)

type People interface {
	Speak(string) string
}

type Stduent struct{}

func TestA(t *testing.T) {

	var peo People = &Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))

}
func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func TestB(t *testing.T) {
	var a a
	defer a.ss("a").ss("b")
	a.ss("c")

}

type a struct {
	m []string
}

func (a *a) ss(s string) *a {
	a.m = append(a.m, s)
	fmt.Println(a)
	return a
}
