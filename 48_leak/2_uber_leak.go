package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func query() int {
	n := rand.Intn(100)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return n
}

// 每次执行此函数，都会导致有两个goroutine处于阻塞状态
func queryAll() int {
	ch := make(chan int) // 无缓冲channel
	go func() { ch <- query() }()
	go func() { ch <- query() }()
	go func() { ch <- query() }()
	<-ch
	<-ch
	return <-ch
}

func main() {
	// 每次循环都会泄漏两个goroutine
	for i := 0; i < 4; i++ {
		queryAll()
		// main()也是一个主groutine
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}

// https://jasonkayzk.github.io/2021/04/21/%E4%BD%BF%E7%94%A8Uber%E5%BC%80%E6%BA%90%E7%9A%84goleak%E5%BA%93%E8%BF%9B%E8%A1%8Cgoroutine%E6%B3%84%E9%9C%B2%E6%A3%80%E6%B5%8B/

// 主要问题发生在 queryAll() 函数里，这个函数在goroutine里往ch里连续三次写入了值，由于这里是无缓冲的ch，所以在写入值的时候，要有在ch有接收者时才可以写入成功，也就是说在从接收者从ch中获取值之前, 前面三个ch<-query() 一直处于阻塞的状态；
//
//当执行到queryAll()函数的 return语句时，ch接收者获取一个值(意思是说三个ch<-query() 中执行最快的那个goroutine写值到ch成功了，还剩下两个执行慢的 ch<-query() 处于阻塞)并返回给调用主函数时，仍有两个ch处于浪费的状态；

// 在Main函数中对于for循环：
//
//第一次：goroutine的总数量为 1个主goroutine + 2个浪费的goroutine = 3；
//第二次：3 + 再个浪费的2个goroutine = 5；
//第三次：5 + 再个浪费的2个goroutine = 7；
//第三次：7 + 再个浪费的2个goroutine = 9；
// 可以看到，主要是ch写入值次数与读取的值的次数不一致导致的有ch一直处于阻塞浪费的状态，我们所以我们只要保存写与读的次数完全一样就可以了；
