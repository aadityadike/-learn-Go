package otherconcepts

import "fmt"

func advanceFunction() {
	mathSquare := square(multiplyItself)
	fmt.Println(mathSquare(5))
}

func multiplyItself(i int) int {
	return i * i
}

func square(anything func(int) int) func(int) int {
	return func(x int) int {
		return anything(x)
	}
}
