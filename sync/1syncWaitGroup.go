package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup

	links := []string{
		"http://www.baidu.com",
		"http://www.jd.com",
		"https://taobao.com",
		"https://www.163.com",
		"http://www.sohu.com",
	}
	for _, link := range links {
		//fmt.Println("url",link)
		waitGroup.Add(1)
		go func(link string) {
			checkLinks(link)
			defer waitGroup.Done()
		}(link)
	}
	waitGroup.Wait()
	//time.Sleep(time.Second*5)
	return
}

func checkLinks(link string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might down")
		return
	}
	fmt.Println(link, "is up!")
}
