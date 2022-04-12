package chanl

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {
	var b chan int
	//go func(b chan int) {
	//	time.Sleep(1 * time.Minute)
	//	b <- 1
	//}(b)
	fmt.Println(<-b)
}
