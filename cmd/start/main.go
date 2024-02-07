package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello, world!")
	var a int = 10
	fmt.Println(a)

	var something bool = true
	fmt.Println(something)

	var name string = "Tanay"
	fmt.Println(utf8.RuneCountInString(name))

	// for int default = 0
	// for strings = ""
	// for bool = false

	// Infer types
	var fullName = "Tanay Vaswani"
	fmt.Println(fullName)

	// Short hand dropping var
	myName := "Tanay Vaswani"
	fmt.Println(myName)

	// someFunc := foo()
	// type unknown until hover
	// define yourself

	const pi float32 = 3.14
	fmt.Println(pi)

	printName(`Tanay`)
}


func printName(name string) {
	fmt.Println(name)
}
