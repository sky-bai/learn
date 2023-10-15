package main

import "fmt"
import "os"

func main() {
	var JAVAHOME string
	JAVAHOME = os.Getenv("PATH")
	fmt.Println("1111", JAVAHOME)

}
