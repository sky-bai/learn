package main

func main() {
	// 右移
	num := 31

	total := 0
	for num > 0 {
		if num&1 == 1 {
			total++
		}
		num = num >> 1
	}

}
