package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

const listenAddres = "localhost:8000"

// normal http
// func main() {
// 	http.HandleFunc("/", handler)
// 	err := http.ListenAndServe(listenAddres, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello, Dear hows your health!")
// }

// websockets

func main() {
	http.Handle("/", websocket.Handler(handler))
	http.ListenAndServe(listenAddres, nil)
}
func handler(c *websocket.Conn) {
	var s string
	fmt.Fscan(c, &s)
	fmt.Println("Received:", s)
	fmt.Fprint(c, "How do you do? Bro1")
}
