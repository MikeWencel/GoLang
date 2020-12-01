package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("this is false")
	case 2 == 4:
		fmt.Println("this is false")
	case 3 == 3:
		fmt.Println("this is true")
	case 4 == 4:
		fmt.Println("this is true, it's should print!")
	}

	switch {
	case false:
		fmt.Println("2 -this is false")
	case 2 == 4:
		fmt.Println("2 -this is false")
	case 3 == 3:
		fmt.Println("2 - this is true")
		fallthrough
	case 4 == 4:
		fmt.Println("2 - this is true, it's should print!")
	default:
		fmt.Println("This is deafault")

	}

}
