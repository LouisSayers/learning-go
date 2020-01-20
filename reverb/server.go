package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:3001")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		text := input.Text()
		if text == "exit" {
			c.Close()
			return
		}
		go echo(c, input.Text(), 1 * time.Second)
	}
}

func echo(c net.Conn, text string, duration time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(text))
	time.Sleep(duration)
	fmt.Fprintln(c, "\t", text)
	time.Sleep(duration)
	fmt.Fprintln(c, "\t", strings.ToLower(text))
}
