package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{"https://google.com", "https://druva.com", "https://facebook.com", "https://youtube.com", "https://weboniselab.com", "https://udemy.com"}

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	// for {
	// 	go checkLink(<-c, c)
	// }

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
	// for l := range c {
	// 	fmt.Println(l)
	// }
}

func checkLink(link string, c chan string) {
	//	time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
