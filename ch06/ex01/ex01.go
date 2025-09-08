package main

import "fmt"

func incr(x *int) {
	*x++
}

func main() {
	x := 10
	p := &x
	incr(p)
	fmt.Println(x)
}
