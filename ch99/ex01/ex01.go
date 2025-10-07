package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var x, y, n int
	fmt.Fscan(in, &x)
	fmt.Fscan(in, &y)
	fmt.Fscan(in, &n)

	for i := 1; i <= n; i++ {
		switch {
		case i%x == 0 && i%y == 0:
			fmt.Println("FizzBuzz")
		case i%x == 0:
			fmt.Println("Fizz")
		case i%y == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
