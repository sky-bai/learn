package main

import (
	"bytes"
	"fmt"
)

func main() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')

	dir1 := path[:sepIndex:sepIndex]
	dir2 := path[sepIndex+1:]

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB
	fmt.Println("cap dir2", cap(dir2))

	dir1 = append(dir1, "suffix"...)

	var n byte = '/'

	path = append(path, n)

	fmt.Println("path =>", string(path)) //prints: path => AAAA/BBBBBBBBB/suffix/

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => uffixBBBB
	fmt.Println("cap dir2", cap(dir2))

}
