package Learnings

import "fmt"

func AdvanceFunction() {
	mathSquare := square(multiplyItself)
	fmt.Println(mathSquare(6))
}

func multiplyItself(i int) int {
	return i * i
}

func square(anything func(int) int) func(int) int {
	return func(x int) int {
		return anything(x)
	}
}
