package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t The inner loop: %d\n", j)
		}
	}

}
