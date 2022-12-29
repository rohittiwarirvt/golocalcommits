package main

import (
	"fmt"
	"os"
)

// panic("a problem")
func mayPanic() {
	panic("a problem")
}
func main() {

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

	// defer ussage
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	//creating a panic
	mayPanic()
	fmt.Println("After mayPanic()")

}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
