package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	str := "cd_admin" + "c071248fa3e9bd2e6d07f7c2b2df87c3" + "d50413ce-8929-17fe-cb6e-e6c75c47e4f2"
	st := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	fmt.Println(st)

	len := len("R2KhtIEcMiYlB9f")
	fmt.Println(len)

}
