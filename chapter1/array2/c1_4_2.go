package main

import "fmt"

func main() {
	a := [2]int{100, 200}
	fmt.Println(a)

	n :=[]int{1,2,3,4,5}
	fmt.Println(n)

	n[2] = 100
	fmt.Println(n)
}