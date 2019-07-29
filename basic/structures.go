package main

import "fmt"

// TestStruct is a test structure
type TestStruct struct {
	field1 int
	field2 string
}

func main() {
	fmt.Println(TestStruct{1, "a"})
	t := TestStruct{}
	t.field1 = 2
	t.field2 = "b"
	fmt.Println(t)
}
