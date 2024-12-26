package main

import (
	"fmt"
	"time"
)

// 在 Go 语言中，可以使用 time 包中的 Format 方法来格式化日期和时间。
//
// 例如，要将当前时间格式化为字符串 "200601021504"，可以使用以下代码：

// 在上面的代码中，我们使用 time.Now() 方法获取当前时间，然后调用 Format 方法将其格式化为指定的字符串。
//
//需要注意的是，Go 语言中的时间格式字符串使用特定的布局标识符，与其他语言可能不同。例如，Go 中的年份布局标识符是 "2006" 而不是 "Y"，月份布局标识符是 "01" 而不是 "m"。因此，在上面的代码中，我们使用的时间格式字符串是 "200601021504" 而不是 "YmdHi"。
//
//完整的代码示例如下：

func main() {
	t := time.Now()
	formatted := t.Format("200601021504")
	fmt.Println(formatted)

	// 年月日时分秒
	// 2021-07-29 15:04:05
	ymdhis := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(ymdhis)

	timeString := time.Now().String()
	fmt.Println(timeString)

	// 年月日
	ymd := time.Now().Format("2006-01-02")
	fmt.Println(ymd)

}
