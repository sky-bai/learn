package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	arr := make([]int, 100000000)
	arr[0] = 1
	go func() {
		for {
			time.Sleep(time.Second)
			log.Print("", 1)
		}
	}()
	http.ListenAndServe(":39090", nil)
}
