package main

import (
	"fmt"
	"testing"
)

func Test1460(t *testing.T) {

	//in {}

	in := []int{-6, 0, 2, 40}
	val, b := DAY1460(in)

	if b {
		fmt.Println("Values at:", val, b)
		return
	}

	t.Errorf("array does fould a value: %d %d %t ", in, val, b)

}

func Test1462(t *testing.T) {
	in := []any{15, 7, 1, 1, "+", "-", "/", 3, "*", 2, 1, 1, "+", "+", "-"}
	res := DAY1462(in)

	if res != 5 {
		t.Errorf("expected 5 got: %d", res)
	}
}

func Test1463(t *testing.T) {
	in := []int{3, 4, 5}
	DAY1463(in)
}
