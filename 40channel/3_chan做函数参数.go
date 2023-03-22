package main

func main() {
	// channel做函数参数
	// 1.定义一个channel
	// 2.把channel作为参数传递给函数

	ch1 := make(chan<- int, 1)
	Chan(ch1)

}

func Chan(ch chan<- int) {

}
