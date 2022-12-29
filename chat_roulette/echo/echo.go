package main

import (
	"io"
	"log"
	"net"
)

var listenAdd string = "localhost:8000"

func main() {
	l, err := net.Listen("tcp", listenAdd)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go io.Copy(c, c)
	}
}
