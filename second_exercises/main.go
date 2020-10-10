package main

import (
	"fmt"
)

func main() {
	//short declaration
	a := 12
	// My longer version
	fmt.Printf("%b\n", a)
	fmt.Printf("%#x\n", a)
	fmt.Printf("%d\n", a)
	//other

	fmt.Printf("%b\t%#x\t%d", a, a, a)

}
