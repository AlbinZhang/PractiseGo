package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	/*
		done := make(chan struct{})
		go func() {
			io.Copy(os.Stdout, conn)
			log.Println("done")
			done <- struct{}{}
		}()

		mustCopy(conn, os.Stdin)
		conn.Close()
		<-done
	*/

	go print(conn)

	fmt.Fprintln(conn, "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA\n")
	input := bufio.NewScanner(conn)
	for input.Scan() {
		log.Println(input.Text())
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Println(err)
	}
}

func print(conn net.Conn) {
	get := bufio.NewScanner(os.Stdin)
	for get.Scan() {
		fmt.Println(get.Text())
		fmt.Fprintln(conn, get.Text())
	}
}
