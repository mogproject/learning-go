package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func main() {
	alice := person{"Alice", 40, "cat"}
	fmt.Println(alice)
}
