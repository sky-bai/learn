package main

import "fmt"

type topicPartition struct {
	topic     string
	partition int32
}

func main() {

	test := make(map[topicPartition][]string)

	test[topicPartition{topic: "test", partition: 1}] = []string{"test1", "test2"}

	t1 := topicPartition{topic: "test", partition: 1}

	if producers, ok := test[t1]; !ok {
		fmt.Println("not ok")
	} else {
		fmt.Println("ok")
		fmt.Println(producers)
	}

	t2 := topicPartition{topic: "test", partition: 1}

	if t1 == t2 {
		fmt.Println("t1 == t2")
	} else {
		fmt.Println("t1 != t2")
	}

}

// 相同值的结构体在map中的key是一样的
