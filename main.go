// Code influenced by: Go. (2026). Functions. GO A Tour of Go. https://go.dev/tour/basics/4
// Demonstration of utilisting a function in Go Language
package main

import "fmt"

// Construct the function at add 2 numbers
func add(a int, b int) int {
	return a + b
}

// Call the function and print results
func main() {
	fmt.Println(add(19, 17))
}
