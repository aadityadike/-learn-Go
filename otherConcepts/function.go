package otherconcepts

import "fmt"

func Function() {
	fmt.Print(Syntaxsugar(2, 3, add))
}

func Syntaxsugar(x, y int, arithmatic func(int, int) int) int {
	return arithmatic(x, y)
}

func add(a, b int) int {
	return a + b
}
