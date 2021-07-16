package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("====start====")
	cmd := exec.Command( "ipconfig")
	read, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	a:=bufio.NewReader(read)
	b,err:=a.ReadString('\n')
	if err!=nil{
		fmt.Println("2:",err)
		return
	}
	fmt.Println(b)
}
