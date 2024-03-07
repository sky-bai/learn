package main

import (
	"github.com/labstack/gommon/log"
	"io"
	"os"
	"os/exec"
)

func main() {
	pr, pw := io.Pipe()
	//  将reader和writer 建立管道 往writer里面写，就可以从reader里面读
	defer pw.Close()

	// tell the command to write to our pipe
	cmd := exec.Command("cat", "fruit.txt")
	cmd.Stdout = pw

	go func() {
		defer pr.Close()
		// copy the data written to the PipeReader via the cmd to stdout
		if _, err := io.Copy(os.Stdout, pr); err != nil {
			log.Fatal(err)
		}
	}()

	// run the command, which writes all output to the PipeWriter
	// which then ends up in the PipeReader
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
