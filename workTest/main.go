package main

import (
	a "execlt1/workTest/A"
	"execlt1/workTest/GOA"
)

func main() {
	//var s = a.A{Name: "aa", Age: 1}
	var q = GOA.GoA{A: a.A{
		Name: "bb",
		Age:  2,
	}}
	q.A.AFunc()
	q.AFunc()
	q.BFunc()

}
