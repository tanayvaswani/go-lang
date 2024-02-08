package main

import (
	"errors"
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

	var numerator int = 10
	var denominator int = 2
	var remainder, result, err = resAndRemain(numerator, denominator)

	if err != nil {
		fmt.Printf(`Error: ` + err.Error())
	} else if remainder == 0 {
		fmt.Println("numerator is completely divisible by denominator")
	}

	fmt.Println(remainder, result)
}


func printName(name string) {
	fmt.Println(name)
}

func resAndRemain(num int, deno int) (int, int, error) {
	var err error

	if deno == 0 {
		err = errors.New("denominator can't be zero")
		return 0, 0, err
	}

	var remainder int = num % deno
	var answer int = num / deno

	return remainder, answer, err
}
