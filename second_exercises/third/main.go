package main

import (
	"fmt"
)

const (
	a     = 42
	b int = 42
)

func main() {
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
}
