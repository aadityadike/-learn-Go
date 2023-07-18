package Learnings

import (
	"fmt"
	"sync"
	"time"
)

func Mutexs() {
	type name struct {
		mux sync.Mutex
	}

	ch := make(chan int, 10)
	oldValue := 1
	for {
		ch <- oldValue
		value := counter1(ch)

		// close(ch)
		fmt.Println(value)
		oldValue = value

		if oldValue > 20 {
			break
		}
	}

}

func counter1(ch chan int) int {
	value := <-ch
	value++
	time.Sleep(time.Millisecond * 100)
	return value
}
