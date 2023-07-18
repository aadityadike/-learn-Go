package Learnings

import (
	"fmt"
)

func Pointers() {
	// value is assigned to a variable.
	a := 3
	// passing the address of value to a function.
	// Print(&a)
	// passing the copy of the value to a function.
	// change(a)
	fmt.Println(&a)
	// printing the value after some operations in function.
	// fmt.Print(a)

	var b *int = &a
	*b = 200
	fmt.Println(b)
	fmt.Println(*b)
}

/*
! '*' goes in front of a variable that holds a memory address and resolves it (it is therefore the counterpart to the & operator). It goes and   gets & the thing that the pointer was pointing at, e.g. *myString.
*/

func Print(value *int) {
	*value = 10
	fmt.Println(value)

}

func change(any int) {
	any = 10
	fmt.Println(&any)
}
