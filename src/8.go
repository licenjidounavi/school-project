  package main

import "fmt"

func main() {
	a := 10
	b := 20
	result := add(a, b)
	fmt.Println("Result:", result)
}

func add(a int, b int) int {
	return a + b
}