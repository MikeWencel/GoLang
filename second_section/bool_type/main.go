package main

import (
	"fmt"
)

var x bool

func main() {
	a := 7
	b := 42
	//You can't compare b or c missed matched types!!!
	// c := "42"
	fmt.Println(x)
	x = true
	fmt.Println(x)
	fmt.Println(a == b)

}
