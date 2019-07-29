package main

import "fmt"

type Info struct {
	name  string
	phone int
}

var id map[int]Info

func main() {
	id = make(map[int]Info)
	id[0] = Info{"Alex", 555}
	id[1] = Info{"Oleg", 111}
	fmt.Println(id)
}
