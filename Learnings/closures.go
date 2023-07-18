package Learnings

import "fmt"

func Closures() {
	cost := counter(email)
	fmt.Printf("This is cost of your emails:%v", cost(2))
	fmt.Printf("This is cost of your emails:%v", cost(2))

}
func email(cost int) int {
	return cost * 2
}

func counter(email func(int) int) func(int) int {
	sum := 0
	return func(i int) int {
		cost := email(i)
		sum += cost
		return sum
	}
}
