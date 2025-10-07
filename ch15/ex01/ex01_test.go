package main

import "testing"

func Test_adder(t *testing.T) {
	result := adder(2, 3)
	if result != 5 {
		t.Error("wrong result: ", result)
	}
}
