package main

import "fmt"

func main() {
	greeting := greet()
	fmt.Println(greeting)
}

// greet returns a greeting to the world
func greet() string {
	return "Hello world"
}
