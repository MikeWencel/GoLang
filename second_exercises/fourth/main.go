package main

import (
	"fmt"
)

func main() {
	a := 12
	fmt.Printf("%b\t%#x\t%d\n", a, a, a)
	b := a << 1
	fmt.Printf("%b\t%#x\t%d", b, b, b)
}
