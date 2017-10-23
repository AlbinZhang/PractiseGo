package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go func() {
		for {

			buf := make([]byte, 100)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(buf[:n]))
		}
	}()

	conn.Write([]byte("Hello world!"))
	write := bufio.NewScanner(os.Stdin)
	for write.Scan() {
		_, err := conn.Write([]byte(write.Text()))
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
