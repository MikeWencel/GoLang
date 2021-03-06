package main

import (
	"fmt"
)

func main() {
	// x :=type // Composite literal
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[0])

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println("********************")
	fmt.Println(x[1])
	fmt.Println(x[0:3])
	fmt.Println("********************")
	for i := 0; i < (len(x)); i++ {
		fmt.Println(x[i])
	}
	fmt.Println("********************")

	x = append(x, 12)
	fmt.Println(x)

	y := []int{202, 303, 13, 12}
	x = append(x, y...)
	fmt.Println(x)
	fmt.Println("********************")
	x = append(x[:0], x[5:]...)
	fmt.Println(x)
	fmt.Println("********************")

}
