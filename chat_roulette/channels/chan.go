package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)

	go fib(ch)
	for i := 0; i <= 2000000; i++ {
		fmt.Println(<-ch)
	}

}

func fib(ch chan int) {
	i, j := 0, 1

	for {
		ch <- j
		i, j = j, j+i
	}

}
