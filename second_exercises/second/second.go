package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println(a)

	fmt.Println(a == a)
	fmt.Println(a <= a)
	fmt.Println(a >= a)
	fmt.Println(a != a)
	fmt.Println(a < a)
	fmt.Println(a > a)
}
