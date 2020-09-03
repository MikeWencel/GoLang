package main

import (
	"fmt"
)

func main() {
	// When it's first time we have to declere with short declatration ':='
	x := 42
	fmt.Println(x)
	x = 100
	fmt.Println(x)
	y := 100 + x
	fmt.Println(y)

}
