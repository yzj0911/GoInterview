package GOA

import (
	a "execlt1/workTest/A"
	"fmt"
)

type GoA struct {
	a.A `inject:""`
}

func (a *GoA) AFunc() {
	fmt.Println("-----a------")
}

func (a *GoA) BFunc() {
	fmt.Println("==========b==========")
}
