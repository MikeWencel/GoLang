package main

import (
	"fmt"
)

var a int

//Creating new type names "HOT DOG" and it's int
type hotdog int

//Creating VARIABLE of type HOT DOG
var b hotdog

func main() {
	a = 82
	b = 73
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
