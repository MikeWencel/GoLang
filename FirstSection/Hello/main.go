package main

import "fmt"

func main() {
	fmt.Println("Hello!!!")

	foo()
	fmt.Println("something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("and the we exited!!!")
}

// Control Flow:
// (1) sequence
// (2) loop; iterative
// (3) conditional
