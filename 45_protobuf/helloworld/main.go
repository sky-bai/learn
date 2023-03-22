package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"learn/45_protobuf/helloworld/types"
)

func main() {
	req := &types.Request{
		Data: "hello",
	}
	encoded, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}

	fmt.Println("---===", encoded)

	var unmarshaledReq types.Request
	err = proto.Unmarshal(encoded, &unmarshaledReq)
	if err != nil {
		panic(err)
	}
	fmt.Println(unmarshaledReq)
}
