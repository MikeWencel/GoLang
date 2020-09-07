package main

import (
	"fmt"
)

type typek int

var x typek

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
