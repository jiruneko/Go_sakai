package main

import "fmt"

func one(x *int) {
	// x = 1
	*x = 1
}

func main() {
	// var n int = 100
	// one(n)
	// fmt.Println(n)
	var n int = 100
	one(&n)
	fmt.Println(n)
}
