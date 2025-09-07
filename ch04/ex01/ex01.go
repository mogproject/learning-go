package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := range 10 {
		fmt.Printf("i=%d, rand=%d\n", i, rand.Intn(10))
	}

	samples := []string{"hello", "ππ", "こんにちは"}
	for _, s := range samples {
		for i, r := range s {
			fmt.Println(i, r, string(r))
		}
	}

outer:
	for i, v := range []int{2, 4, 6} {
		for j, u := range []int{5, 7, 9, 7, 5} {
			if (v+u)%3 == 0 {
				continue outer
			}
			fmt.Printf("i=%d, j=%d, v=%d, u=%d\n", i, j, v, u)
		}
	}
}
