package main

import "fmt"

func main() {

	for i := 0; i < 4; i++ {
		if i == 3 {
			for a := 0; a < 5; a++ {
				fmt.Println("====")
				if a==2{
					goto A
				}
			}
		}
	A:
		fmt.Println("end")
	}

}
