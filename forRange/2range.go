package main

func main() {
	slice := []int{0, 1, 2, 3}
	myMap := make(map[int]*int)

	for key, value := range slice {
		tem := value
		myMap[key] = &tem
	}
}

// 因为 for range 遍历出来的value是同一块地址 所以这里取value的地址，他的值都是最后一个值
