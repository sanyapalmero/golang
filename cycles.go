package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(i, sum)
	}

	sum2 := 0
	for sum2 < 10 { //like while
		sum2 += 1
	}
	fmt.Println(sum2)

}
