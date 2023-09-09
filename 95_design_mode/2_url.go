package main

import "fmt"

func main() {
	taskIdAndUrlMedia := make(map[string]string)
	reply := make(map[string]string)
	if url, ok := taskIdAndUrlMedia["1"]; ok {
		reply["attach_url"] = url
	} else {
		reply["attach_url"] = url
		reply["attach_url111"] = ""
		fmt.Println("11111", reply)
	}
	fmt.Println("-----", reply)
}
