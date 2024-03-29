package main

import (
	"fmt"
	"log"

	"github.com/kevwan/mapreduce/v2"
)

// 流数据是一组顺序、大量、快速、连续到达的数据序列,一般情况下，流数据可被视为一个随时间延续而无限增长的动态数据集合。
// 但实际业务场景中多个依赖如果有一个出错我们期望能立即返回而不是等所有依赖都执行完再返回结果，而且WaitGroup中对变量的赋值往往需要加锁，每个依赖函数都需要添加Add和Done对于新手来说比较容易出错。
// commit
// 并行求平方和
func main() {
	// source 数据源 val 最终获得的值
	val, err := mapreduce.MapReduce(func(source chan<- int) {
		// generator
		for i := 0; i < 10; i++ {
			source <- i
		}
	}, func(i int, writer mapreduce.Writer[int], cancel func(error)) {
		// mapper
		writer.Write(i * i)
	}, func(pipe <-chan int, writer mapreduce.Writer[int], cancel func(error)) {
		// reducer
		var sum int
		for i := range pipe {
			sum += i
		}
		writer.Write(sum)
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result:", val)
}
