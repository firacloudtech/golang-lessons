//Your first Go application

// 1. Create a file named main.go in your root folder. Declare your package name here.
package main

import control "github.com/yogenp/golang-lessons/internals/2.control"

// 2. import packages here. If you are importing an internal package, make sure to run `go mod tidy`

//3. Declare your function here

func main() {

	// call a function from the package
	// basics.Basics()
	// base.Base()
	//variables.PrintVariables()
	control.Control()

}

// 4. Now, run `go run main.go` in your terminal
// 5. To compile your program, run `go build main.go`. Then, run `./main`.
