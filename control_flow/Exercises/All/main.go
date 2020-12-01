package main

import (
	"fmt"
)

func main() {
	fmt.Println("***************")
	fmt.Println("Exercise numer 7")
	fmt.Println("***************")
	x := "James Bond"
	y := 7

	if x == "James Bond" {
		fmt.Println("This is " + x)
	} else if x == "James Bond" && y == 7 {
		fmt.Println("James Bond is 007 agent!")
	} else {
		fmt.Println("Who is this guy?")
	}
	fmt.Println("***************")
	fmt.Println("Exercise numer 8")
	fmt.Println("***************")
	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("Should print")
	}

	fmt.Println("***************")
	fmt.Println("Exercise numer 9")
	fmt.Println("***************")
	favSport := "kitesurfing"

	switch favSport {
	case "football":
		fmt.Println("Go to the pitch")
	case "tenis":
		fmt.Println("Go on court")
	case "kitesurfing":
		fmt.Println("Go to Rewa")
	case "aviation":
		fmt.Println("Go to the airpoirt")
	}
	fmt.Println("***************")
	fmt.Println("Exercise numer 10")
	fmt.Println("***************")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
