package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	fmt.Println("Int:", x)
	fmt.Println("String:", y)
	fmt.Println("Bool:", z)

	s := fmt.Sprintf("%d is Int, %s is String, %t is Bool\n", x, y, z)
	fmt.Println(s)
}
