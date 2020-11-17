package main

import (
	"fmt"
)

func main() {
	x := 40

	if x == 40 {
		fmt.Println("Our value was 40")
	} else if x == 41 {
		fmt.Println("our value was 41")
	} else {
		fmt.Println("our value was not 40")
	}
}
