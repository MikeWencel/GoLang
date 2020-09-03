package main

import "fmt"

// Variable can be declare outside of func body
var y = 43

//DECLARE variable of type INT with IDENTIFIER "z", value assigned is "0"
var z int

func main() {
	//short declaration operator is working only in function body
	// Declare Var and Assign a Value
	x := 99
	fmt.Println(x)

	fmt.Println(y)
	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println("again: ", y)
}
