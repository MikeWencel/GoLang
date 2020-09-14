package main

import (
	"fmt"
)

var x int
var y float64
var z byte

func main() {
	s := "Michał i Kasia Wencel"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	z := 102
	bz := string(z)
	fmt.Println(bz)
	fmt.Printf("%T\n", bz)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
}
