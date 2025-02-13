package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var Up = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := Up.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("recv: %s", message)
		fmt.Printf("mt: %d", mt)

		//err = conn.WriteMessage(mt, message)
		//if err != nil {
		//	fmt.Println(err)
		//	break
		//}
	}

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}
