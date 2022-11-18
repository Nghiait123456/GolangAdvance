package main

import (
	"fmt"
	"net"
)

func handler(c net.Conn) {
	fmt.Println("start handle one request")
	fmt.Println("end job")
	c.Write([]byte("ok"))
	c.Close()
	return
}

func main() {
	fmt.Println("Listening on :8000. Send something using nc: echo hello | nc localhost 8000")
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	fmt.Println("start listen process")
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c)
	}
}
