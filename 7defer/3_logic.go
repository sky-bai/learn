package main

func main() {
	normalReturn := false // 先定义没有成功

	// 那么成功了就设置成true
	//normalReturn = true
	if normalReturn { //
		println("normal return")
	}

	if !normalReturn {
		println("no normal return")
	}

}
