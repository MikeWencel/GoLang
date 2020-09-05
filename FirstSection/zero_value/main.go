package main

import (
	"fmt"
)

var word string
var number int

// TO NIE BĘDZIE DZIAŁAĆ - BRAK TYPU

// var y
var name string

var num = 911

func main() {
	fmt.Println("Value of string 'z' is:", word, ". - It's empty, but it's declared")
	fmt.Printf("%T\n", word)
	// fmt.Printf("%T", y)
	fmt.Printf("%T\n", number)
	fmt.Printf("%T\n", num)
	// %v the value in a default format
	fmt.Printf("%v\n", num)
	//Go-Sytax representation of the value
	fmt.Printf("%#v\n", num)
	//binary of num = 9
	fmt.Printf("%b\n", num)
	//hexadecimal of num = 9
	fmt.Printf("%x\n", num)
	//hexadecimal of num = 9 with 0 in front
	fmt.Printf("%#x\n", num)

	fmt.Printf("%#x\t%b\t%x", num, num, num)

	// \a	Alert or bell
	// \b	Backspace
	// \\	Backslash
	// \t	Horizontal tab
	// \n	Line feed or newline
	// \f	Form feed
	// \r	Carriage return
	// \v	Vertical tab
	// \'	Single quote (only in rune literals)
	// \"	Double quote (only in string literals)
}
