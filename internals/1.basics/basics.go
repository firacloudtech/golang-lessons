package basics

import (
	"fmt"
)

func Basics() {

	// 1. How to print to terminal?
	fmt.Println("Hello world!")
	fmt.Printf("Hello %v", "World\n")

	// 2. How to declare variables?
	var a = "Hello"
	b := 2
	const c string = "10"

	fmt.Printf("String:  %v \n", a)
	fmt.Printf("Shorthand declaration:  %v \n", b)
	fmt.Printf("Constant: %v \n", c)

}
