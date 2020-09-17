package main

import (
	"fmt"
)

const (
	_ = iota
	// kb = 1024
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {

	fmt.Printf("%d\t\t\t%bb\n", kb, kb)
	fmt.Printf("%d\t\t\t%bb\n", mb, mb)
	fmt.Printf("%d\t\t%bb\n", gb, gb)
}
