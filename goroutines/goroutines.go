package main

import (
	"fmt"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	r := make(chan int32)
	f("direct")

	go f("goroutine")

	// go func(msg string) {
	// 	time.Sleep(3 * time.Second)
	// 	r <- 2
	// 	fmt.Println(msg)
	// }("going")

	//time.Sleep(time.Second)

	fmt.Println(<-r)
	fmt.Println("done")
}
