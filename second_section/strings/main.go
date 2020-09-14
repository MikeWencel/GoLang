package main

import (
	"fmt"
)

var x int
var y float64

func main() {
	s := "Michał i Kasia Wencel"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
}
