package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var i, j int
	i = 1
	marshal, err := json.Marshal(i)
	if err != nil {
		panic(err)
		return
	}
	err = json.Unmarshal(marshal, &j)
	if err != nil {
		panic(err)
		return
	}
	fmt.Print("---", j)

	k := 1
	err = json.Unmarshal([]byte(k), &j)
	if err != nil {
		panic(err)
		return
	}
	fmt.Print("---", j)
}
