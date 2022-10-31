package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type PayLoad struct {
	Content string
}

func main() {
	pr, pw := io.Pipe()

	go func() {
		// close the writer, so the reader knows there's no more data
		defer pw.Close()

		// write json data to the PipeReader through the PipeWriter
		if err := json.NewEncoder(pw).Encode(&PayLoad{Content: "Hello Pipe!"}); err != nil {
			log.Fatal(err)
		}
	}()

	// JSON from the PipeWriter lands in the PipeReader
	// ...and we send it off...
	if _, err := http.Post("http://example.com", "application/json", pr); err != nil {
		log.Fatal(err)
	}
}
