/**
 * @author [zhangfeng]
 * @email [zhangfeng1@hubcloud.com.cn]
 * @create date 2017-10-31 07:37:31
 * @modify date 2017-10-31 07:37:31
 * @desc [description]
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	conn *websocket.Conn
	send chan []byte
}

func serveWs(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	Write(conn)
}

func Write(conn *websocket.Conn) {
	for i := 10; i < 100; i++ {
		conn.SetWriteDeadline(time.Now().Add(writeWait))
		w, err := conn.NextWriter(websocket.TextMessage)
		if err != nil {
			return
		}

		w.Write([]byte(fmt.Sprintln(i)))
		w.Close()
	}
}
