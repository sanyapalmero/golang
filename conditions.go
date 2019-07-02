package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := 5
	b := rand.Intn(10)

	if a < b {
		fmt.Println(a, b, "b")
	} else {
		fmt.Println(a, b, "a")
	}
}
