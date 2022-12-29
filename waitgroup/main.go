package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	go func() {
		wg.Add(1)
		defer wg.Done()
		getAndPrintData("https://boot.dev")
	}()
	go func() {
		wg.Add(1)
		defer wg.Done()
		getAndPrintData("https://github.com")
	}()
	go func() {
		wg.Add(1)
		defer wg.Done()
		getAndPrintData("https://gitlab.io")
	}()

	//wg.Wait()
}

func getAndPrintData(url string) {
	resp, _ := http.Get(url)
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Println(buf.String())
}
