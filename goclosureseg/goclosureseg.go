package main

import "fmt"

func getInt() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	myNewIntR := getInt()
	fmt.Println(myNewIntR())
	fmt.Println(myNewIntR())
	fmt.Println(myNewIntR())
	fmt.Println(myNewIntR())

	myNewIntS := getInt()
	fmt.Println(myNewIntS())
}
