package main

import "fmt"

func main() {
	fmt.Println("this is int type -> 1 + 1 =", 1+2)                    // adding two ints 1 and 2
	fmt.Println("this is floating point type -> 1.0 + 2.0 =", 1.0+1.0) // adding two float points 1.0 and 2.0

	fmt.Println(len("Hello, gopher!"))            // string opr, get len of string
	fmt.Println("thisIsString"[5])                // string opr, access any char of string
	fmt.Println("Hello, " + "world -" + "gopher") // concate the multiple strings to crate final string

	// boolean opr
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)

}
