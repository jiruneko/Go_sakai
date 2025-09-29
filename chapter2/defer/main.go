package main

import "fmt"

func foo() {
	defer fmt.Println("World foo")
	fmt.Println("Hello foo")
}
func main() {
	foo()
	defer fmt.Println("world")
	fmt.Println("hello")
}
