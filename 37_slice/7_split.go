package main

import (
	"fmt"
	"strings"
)

func main() {

	arr := strings.Split("#1:803278221433890:1:*,00000002,XT,true+广东省+深圳市+龙岗区,A,170704,090E3B,FF1035E3,00FD7DC0,0059,7594,000000010000,5A,4,000064,3#", ",")
	fmt.Println(len(arr))

}
