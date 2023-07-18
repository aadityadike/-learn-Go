package Learnings

import (
	"fmt"
	"time"
)

func Channels() {
	var i int = 10
	concurrentFib(i)
}

func concurrentFib(i int) {
	ch := make(chan int)

	go func() {
		fibonacci(i, ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
	_, ok := <-ch
	fmt.Print(ok)

}

func fibonacci(num int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < num; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}
