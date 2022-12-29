package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// normat await
	r := <-longRunningTask()
	fmt.Println(r)

	// promise.all
	aCh, bCh, cCh := longRunningTask(), longRunningTask1(), longRunningTask()

	a, b, c := <-aCh, <-bCh, <-cCh
	fmt.Println(a, b, c)

	// promice.race
	var v int32
	select {
	case v = <-one():
	case v = <-two():
	}
	fmt.Println("First Resolved by ")
	fmt.Println(v)
}

func longRunningTask() <-chan int {
	r := make(chan int)

	go func() {
		defer close(r)
		time.Sleep(time.Second * 3)
		r <- rand.Int()
	}()

	return r
}

func longRunningTask1() <-chan int {
	r := make(chan int)

	go func() {
		defer close(r)

		time.Sleep(time.Second * 3)
		panic("hello no")
		r <- rand.Int()
	}()

	return r
}

func one() <-chan int32 {
	r := make(chan int32)
	go func() {
		time.Sleep(time.Second * time.Duration(rand.Int63n(3000)))
		r <- 1
	}()

	return r
}

func two() <-chan int32 {
	r := make(chan int32)
	go func() {

		time.Sleep(time.Millisecond * time.Duration(rand.Int63n(1000)))
		//	time.Sleep(time.Millisecond * time.Duration(rand.Int63n(1000)))
		r <- 2
	}()

	return r

}
