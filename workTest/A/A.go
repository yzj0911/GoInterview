package A

import "fmt"

type A struct {
	Name string
	Age  int
}

func (a *A) AFunc() {
	fmt.Println(a.Name)
}
func (a *A) BFunc() {
	fmt.Println(a.Age)
}
