package main

import (
	"fmt"
	"sync"
	"time"
)

type Promise struct {
	wg  sync.WaitGroup
	res string
	err error
}

func NewPromise(f func() (string, error)) *Promise {
	p := &Promise{}
	p.wg.Add(1)
	go func() {
		p.res, p.err = f()
		p.wg.Done()
	}()
	return p
}

func (p *Promise) Then(r func(string), e func(error)) {
	go func() {
		p.wg.Wait()
		if p.err != nil {
			e(p.err)
			return
		}
		r(p.res)
	}()
}

func exampleTicker() (string, error) {
	<-time.Tick(time.Second * 1)
	return "hi", nil
}

func main() {
	doneChan := make(chan int)
	var p = NewPromise(exampleTicker)
	p.
		Then(func(result string) { fmt.Println(result); doneChan <- 1 }, func(err error) { fmt.Println(err) })
	<-doneChan
}
