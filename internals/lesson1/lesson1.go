//Your first Go application

// 1. Create a file named main.go in your root folder. Declare your package name here.
package lesson1

// 2. import packages here

import "fmt"

//3. Declare your function here

func PrintHelloWorld() {

	fmt.Printf("Hello %v\n", "World")

}

// 4. Now, run `go run main.go` in your terminal
// 5. To compile your program, run `go build main.go`. Then, run `./main`.
