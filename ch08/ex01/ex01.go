package main

import "fmt"

type Thinger interface {
	Thing()
}

type ThingerInt int

func (t ThingerInt) Thing() {
	fmt.Println("ThingInt:", t)
}

type ThingerSlice []int

func (t ThingerSlice) Thing() {
	fmt.Println("ThingSlice:", t)
}

func Comparer[T comparable](t1, t2 T) {
	if t1 == t2 {
		fmt.Println("Equal!!!")
	}
}

func main() {
	a := 10
	b := 10
	Comparer(a, b)

	var a2 ThingerInt = 20
	var b2 ThingerInt = 20
	Comparer(a2, b2)

	var a3 ThingerSlice = []int{1, 2, 3}
	var b3 ThingerSlice = []int{1, 2, 3}
	// Comparer(a3, b3)

	var a4 Thinger = a2
	var b4 Thinger = b2
	Comparer(a4, b4)

	a4 = a3
	b4 = b3
	Comparer(a4, b4)
}
