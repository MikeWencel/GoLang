package main

import (
	"fmt"
)

func main() {
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("Done")

	for {
		if x > 20 {
			break
		}
		fmt.Println(x)
		x++
	}
}
