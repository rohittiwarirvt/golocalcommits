package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	dataChan := make(chan int)

	// go func() {
	// 	dataChan <- 213
	// }()

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("Hi bro!")
	// 	fmt.Println(<-dataChan)
	// 	//dataChan <- 213
	// 	fmt.Println("Hi bro!")
	// 	//fmt.Println(<-dataChan)

	// }()
	// dataChan <- 213

	// time.Sleep(3 * time.Second)
	// fmt.Println(<-dataChan)

	wg := sync.WaitGroup{}
	go func() {
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := doWork()
				dataChan <- result
			}()

		}
		wg.Wait()
		close(dataChan)
	}()

	for val := range dataChan {
		fmt.Println(val)
	}
}

func doWork() int {
	time.Sleep(1 * time.Second)
	return rand.Intn(100)
}
