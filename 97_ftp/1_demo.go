package main

import (
	"github.com/jlaffaye/ftp"
	"github.com/robfig/cron/v3"
	"io"
	"log"
	"os"
	"time"
)

func main() {

	c := cron.New(
		cron.WithLogger(
			cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))

	c.AddFunc("@every 2h", func() {
		TimeToUpdateFtpJob()
	})

	c.Start()
	select {}

}

func TimeToUpdateFtpJob() {
	c, err := ftp.Dial("ftp.example.org:21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("anonymous", "anonymous")
	if err != nil {
		log.Fatal(err)
	}

	// Do something with the FTP conn

	r, err := c.Retr("test-file.txt")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	buf, err := io.ReadAll(r)
	println(string(buf))

	// todo download to file

	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}
