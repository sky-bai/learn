package main

import "path/filepath"

func main() {

	objectName := "cf_ori/20243/gaoDe/186334436368697/1710788747727.jpg"
	dir, file := filepath.Split(objectName)
	println(dir)
	println(file)
}
