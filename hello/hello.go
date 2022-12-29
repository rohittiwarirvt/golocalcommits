package main

import (
	"fmt"
	"time"
)

var test string
var lang = "go"

//lang = "go"

func main() {

	// variables and compiles
	name := "Druva"
	lang = "go"

	var b, c int = 1, 2
	fmt.Println(b, c)

	const S = "constant"
	//var test string
	//test = "test variable"
	fmt.Println("Lets learn a bit about ", lang)
	fmt.Println("Hi Welcome to Developer forum", name)

	// looops
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// no do while, while loop

	// if else
	//this statement are available in all branches.

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// switch
	t := time.Now()
	switch t {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}
