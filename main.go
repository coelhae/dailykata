package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println("daily kata stuff!")
}

// A fixed point in an array is an element whose value is equal to its index.
// Given a sorted array of distinct elements, return a fixed point, if one exists. Otherwise, return False.
// For example, given [-6, 0, 2, 40], you should return 2. Given [1, 5, 7, 8], you should return False.
func DAY1460(in []int) (int, bool) {

	i := 0
	for range in {
		if in[i] == i {
			return i, true
		}
		i++
	}

	return -1, false
}

// Given an arithmetic expression in Reverse Polish Notation, write a program to evaluate it.
//The expression is given as a list of numbers and operands. For example: [5, 3, '+'] should return 5 + 3 = 8.
//For example, [15, 7, 1, 1, '+', '-', '/', 3, '*', 2, 1, 1, '+', '+', '-'] should return 5, since it is equivalent to ((15 / (7 - (1 + 1))) * 3) - (2 + (1 + 1)) = 5.

type stack []int

func (s *stack) push(n int) {
	*s = append(*s, n)
}

func (s *stack) pop() (int, bool) {
	size := len(*s)
	if size > 0 {
		out := (*s)[size-1]
		*s = (*s)[:size-1]
		return out, true
	}

	return 0, false
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}
func div(x, y int) int {
	return x / y
}
func mult(x, y int) int {
	return x * y
}

func DAY1462(in []any) int {

	var s = make(stack, 0) // bad
	// todo: all empty validation

	//var result int = 0

	//index := len(in)
	for _, v := range in {
		if reflect.TypeOf(v).Name() == "int" {
			var inteiro interface{} = v
			var x int = inteiro.(int)
			s.push(x)
		} else {

			var st interface{} = v
			switch st.(string) {
			case "+":
				w, _ := s.pop()
				z, _ := s.pop()
				s.push(add(w, z))
			case "-":
				w, _ := s.pop()
				z, _ := s.pop()
				s.push(sub(z, w))
			case "*":
				w, _ := s.pop()
				z, _ := s.pop()
				s.push(mult(z, w))
			case "/":
				w, _ := s.pop()
				z, _ := s.pop()
				s.push(div(z, w))
			}

		}
	}
	fmt.Println("result :", s)
	out, _ := s.pop()
	return out
}

/* -----------------------------------------------------------------*/

//Given an array of integers, determine whether it contains a Pythagorean triplet.
//Recall that a Pythogorean triplet (a, b, c) is defined by the equation a2+ b2= c2.
func DAY1463(in []int) bool {
	//TODO: missing minimum and max index validations. assume array is ok.

	a := math.Pow(float64(in[0]), 2)
	b := math.Pow(float64(in[1]), 2)
	c := math.Pow(float64(in[2]), 2)

	// not all permutations are done. this assumes just that order of validation is done
	if a+b == c {
		fmt.Println(c)
		return true
	}

	return false
}

/* -----------------------------------------------------------------*/
