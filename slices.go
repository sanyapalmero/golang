package main

import "fmt"

func main() {
	array := [6]int{1, 2, 3, 4, 5, 6}

	var arraySlice []int = array[0:3]
	fmt.Println(arraySlice)

	fmt.Println("arraySlice len:", len(arraySlice))
	fmt.Println("arraySlice cap:", cap(arraySlice))

	//create slice as array
	q := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(q)

	//make
	m := make([]int, 5)
	fmt.Println(m)

	m = append(m, 5, 5, 5)
	fmt.Println(m)

	for i, elem := range m {
		fmt.Printf("index: %d", i)
		fmt.Printf(" elem: %d\n", elem)
	}
}
