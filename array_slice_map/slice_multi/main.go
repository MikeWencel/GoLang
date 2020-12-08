package main

import (
	"fmt"
)

func main() {
	jb := []string{"Mike", "Wen", "Todd"}
	fmt.Println(jb)
	mp := []string{"Mike", "Wen", "Todd"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
