package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("first", i)
		for j := 0; j < 10; j++ {
			fmt.Println("second", j)
		}
	}

}
