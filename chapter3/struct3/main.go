package main

import "fmt"

type Vertex struct {
	X int
	Y int
	S string
}

func main() {
	v2 := Vertex{X: 1}
	fmt.Println(v2)
	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Println(v4)
	var v5 Vertex
	fmt.Println(v5)

	v6 := new(Vertex)
	fmt.Println(v6)
}
