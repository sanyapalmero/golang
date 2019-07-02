package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func add(a int, b int) int {
	return a + b
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(add(5, 7))

	test := "test"
	fmt.Println(test)
}
