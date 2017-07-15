package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)

	}
}

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	message  = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			//clients[cli] = false
			delete(clients, cli)
			close(cli)
		case msg := <-message:
			fmt.Println("当前人数(", len(clients), ")", msg)
			for cli := range clients {
				cli <- msg
			}
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()

	entering <- ch
	message <- who + " entering."

	/*
		input := bufio.NewScanner(conn)
		for input.Scan() {
		}
	*/

	time.Sleep(1 * time.Second)
	for {
		buf := make([]byte, 50)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		message <- who + " " + string(buf[:n]) //+ "\n"
	}

	leaving <- ch
	message <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		_, err := conn.Write([]byte(msg))
		if err != nil {
			log.Println("clientWriter error:", err)
		}
	}
}
