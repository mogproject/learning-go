package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func doSomeWork() int {
	x := rand.Intn(10)
	fmt.Println("x:", x)
	if x != 9 {
		return x
	} else {
		time.Sleep(3 * time.Second)
		return 100
	}
}

func timeLimit[T any](worker func() T, limit time.Duration) (T, error) {
	out := make(chan T, 1)
	ctx, cancel := context.WithTimeout(context.Background(), limit)
	defer cancel()

	go func() {
		out <- worker()
	}()

	select {
	case result := <-out:
		return result, nil
	case <-ctx.Done():
		var zero T
		return zero, errors.New("time limit exceeded")
	}
}

func main() {
	for range 20 {
		result, err := timeLimit(doSomeWork, 2*time.Second)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			fmt.Printf("Result: %v\n", result)
		}
	}
}
