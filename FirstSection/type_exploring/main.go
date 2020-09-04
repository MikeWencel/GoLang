package main

import (
	"fmt"
)

var y = 99

// We DECLARE that Z is STRING
// We cannot change to INT
// THIS IS STATIC PROGRAMMING LANGUAGE NOT DYNAMIC!
var z = "Few words"

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Println("%T\n", z)

}
