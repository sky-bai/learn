package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []string{"hello", "", "world", "yes", "hello", "nihao"}
	//sort.Strings(a)
	fmt.Println(a)
	b := RemoveDuplicatesAndEmpty(a)
	fmt.Println(len(b))
	fmt.Println(b)

}
func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	sort.Strings(a)
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if i > 0 && a[i-1] == a[i] {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}
