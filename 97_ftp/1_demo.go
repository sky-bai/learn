package main

import (
	"bytes"
	"fmt"
	"github.com/jlaffaye/ftp"
	"github.com/robfig/cron/v3"
	"io"
	"log"
	"sync"
	"time"
)

const OneSecond = 1*time.Second + 50*time.Millisecond

type syncWriter struct {
	wr bytes.Buffer
	m  sync.Mutex
}

func (sw *syncWriter) Write(data []byte) (n int, err error) {
	sw.m.Lock()
	n, err = sw.wr.Write(data)
	sw.m.Unlock()
	return
}

func (sw *syncWriter) String() string {
	sw.m.Lock()
	defer sw.m.Unlock()
	return sw.wr.String()
}

func newBufLogger(sw *syncWriter) cron.Logger {
	return cron.PrintfLogger(log.New(sw, "", log.LstdFlags))
}

func main() {
	c := cron.New(
		cron.WithChain(cron.Recover(cron.DefaultLogger)))
	c.Start()

	c.AddFunc("@every 5s", func() {
		fmt.Println("---")
	})

	c.Start()
	select {}

}

func TimeToUpdateFtpJob() {
	c, err := ftp.Dial("ftp://agnss.allystar.com", ftp.DialWithTimeout(5*time.Second))
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

// 定时任务的recover 错误日志
// 任务panic的recover
// 任务执行的日志
