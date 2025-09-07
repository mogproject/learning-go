package main

import "fmt"

func main() {
	var x []int
	x = append(x, 10, 20)
	fmt.Println(x)

	y := make([]int, 5)
	fmt.Println(y)

	z := []int{}
	z = append(z, 15)
	z = append(z, 5)
	fmt.Println(z)

	m := map[string]int{}
	m["abc"] = 3
	m["de"] = 4
	m[""] = 5
	fmt.Println(m)

	_, ok := m["ab"]
	fmt.Println(ok) // false

	_, ok = m["abc"]
	fmt.Println(ok) // true

	delete(m, "abc")
	_, ok = m["abc"]
	fmt.Println(ok) // false

}
