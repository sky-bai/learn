package main

import (
	"github.com/gorilla/websocket"
	"log"
)

func main() {
	dl := websocket.Dialer{}
	conn, _, err := dl.Dial("ws://127.0.0.1:8888", nil)
	if err != nil {
		log.Println("dl.Dial error: ", err)
		return
	}
	defer conn.Close()

	conn.WriteMessage(websocket.TextMessage, []byte("hello"))
}
