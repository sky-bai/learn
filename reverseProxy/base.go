package main

import (
	"fmt"
	"log"
	"net/http"
)

func hander(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, r)
}

func main() {
	addr := "127.0.0.1:2003"
	http.HandleFunc("/base/dir", hander)
	log.Println("starting server at " + addr)
	http.ListenAndServe(addr, nil)

}
