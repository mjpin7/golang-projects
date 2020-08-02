package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.NewTimer(3 * time.Second)

problem:
	for i := 0; i < 20; i++ {
		select {
		case <-t1.C:
			break problem
		default:
			fmt.Printf("%d\n", i)
			time.Sleep(1 * time.Second)
		}
	}

	fmt.Printf("test\n")

}
