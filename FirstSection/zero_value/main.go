package main

import (
	"fmt"
)

var z string

func main() {
	fmt.Println("Value of string 'z' is:", z, ". - It's empty, but it's declared")
	fmt.Printf("%T", z)
}
