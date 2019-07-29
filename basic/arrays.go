package main

import "fmt"

func main() {
	var array1 [2]int
	array1[0] = 4
	array1[1] = 5
	fmt.Println(array1)

	var array2 [2]interface{}
	array2[0] = 1
	array2[1] = "s"
	fmt.Println(array2)
}
