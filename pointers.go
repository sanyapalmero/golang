package main

import "fmt"

func main() {
	a := 1
	b := 2

	pointToA := &a
	pointToB := &b

	fmt.Println("point to a: ", pointToA)
	fmt.Println("point to b: ", pointToB)

	fmt.Println("read a through the pointer: ", *pointToA)
	fmt.Println("read b through the pointer: ", *pointToB)

	*pointToA = 3
	*pointToB = 4
	fmt.Println("read new value a through the pointer", *pointToA)
	fmt.Println("read new value b through the pointer", *pointToB)
}
