package main

import "fmt"

func main() {
	var p *int = new(int)
	fmt.Println(p)

	var p2 *int
	fmt.Println(p2)

	var p3 *int = new(int)
	fmt.Println(*p3)
	*p3++
	fmt.Println(*p3)

	// var p4 *int
	// fmt.Println(p4)
	// *p4++
	// fmt.Println(p4)
}
