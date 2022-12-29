package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
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
		go match(c, time.Now())
	}
}

var partner = make(chan io.ReadWriteCloser)

func match(c io.ReadWriteCloser, t time.Time) {
	fmt.Println(c, "Waiting for a partner...")
	select {
	case partner <- c:
		fmt.Println("I am executing bRo!")
		fmt.Println(t.Date())
	case p := <-partner:
		fmt.Println("I am chatting bRo!")
		fmt.Println(t.Date())
		chat(p, c)
	}
}

func chat(a, b io.ReadWriteCloser) {
	fmt.Println(a, "Found one! Say hi.")
	fmt.Println(b, "Found one! Say hi.")
	errc := make(chan error, 1)
	go copy(a, b, errc)
	go copy(b, a, errc)
	if err := <-errc; err != nil {
		log.Println(err)
	}
	a.Close()
	b.Close()
}

func copy(w io.Writer, r io.Reader, errc chan<- error) {
	_, err := io.Copy(w, r)
	errc <- err
}
