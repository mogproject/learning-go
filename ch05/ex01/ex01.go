package main

import "fmt"

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }

type opFuncType func(int, int) int

var opMap = map[string]opFuncType{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6))

	a := []int{4, 3}
	fmt.Println(addTo(3, a...))

	fmt.Println(opMap["+"](3, 12))

	// anonymous function
	f := func(j int) bool {
		for i := 2; i*i <= j; i++ {
			if j%i == 0 {
				return false
			}
		}
		return j >= 2
	}

	for i := range 20 {
		fmt.Printf("i=%d, isPrime=%t\n", i, f(i))
	}
}
