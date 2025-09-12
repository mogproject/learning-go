package main

import "fmt"

func doBusinessLogic(x int) int {
	return x * x
}

func main() {
	in := make(chan int, 5)
	in2 := make(chan int, 5)
	out := make(chan int, 5)

	go func() {
		for {
			select {
			case val := <-in:
				out <- doBusinessLogic(val)
			case val := <-in2:
				out <- doBusinessLogic(val)
			}
		}
	}()

	vals := []int{1, 2, 3, 4, 5}
	for _, v := range vals {
		in <- v * 10
		in2 <- v
	}

	i := 0
	for v := range out {
		fmt.Println(v)
		i++
		if 2*len(vals) <= i {
			break
		}
	}

	close(out)
	close(in)
}
