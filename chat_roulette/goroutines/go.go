package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Second)
	go say("Hwlow", 1)
	go say("This is new worl", 2)
	go say("Thji is another world", 3)
	time.Sleep(5 * time.Second)
}

func say(str string, secs int32) {
	time.Sleep(time.Duration(secs) * time.Second)
	fmt.Println(str)
}
