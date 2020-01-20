package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:3001")
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan bool)

	go func() {
		mustCopy(conn, os.Stdout)
		fmt.Println("DONE!")
		ch <- true
	}()
	mustCopy(os.Stdin, conn)
	conn.Close()
	<-ch
}

func mustCopy(from io.Reader, to io.Writer) {
	if _, err := io.Copy(to, from); err != nil {
		log.Fatal(err)
	}
}
