package main

import (
	"fmt"
)

type Package struct {
	BagStartDate string `json:"bagStartDate"`
}

func main() {

	// 可以使用字符串直接比较的方法来比较日期字符串的大小，但前提是日期字符串的格式是ISO 8601 或类似的格式，
	// 例如 "YYYY-MM-DD HH:mm:ss"。在这种格式下，字符串的字典序（lexicographical order）与时间顺序一致，因此可以直接用 < 和 > 进行比较。
	//
	// 原理解释：
	// 日期字符串格式 "YYYY-MM-DD HH:mm:ss" 是按年、月、日、时、分、秒从高到低排序的，因此比较字符串的字典序时，其结果与时间顺序一致。
	// 如果日期字符串的格式被破坏（例如 "MM-DD-YYYY" 或没有固定长度），直接比较可能会出错。
	// 示例数据
	p1 := Package{BagStartDate: "2024-06-11 11:17:42"}
	p2 := Package{BagStartDate: "2025-06-11 11:17:42"}

	// 比较时间

	if p1.BagStartDate < p2.BagStartDate {
		fmt.Println("p1 的日期更早")
	} else if p1.BagStartDate > p2.BagStartDate {
		fmt.Println("p2 的日期更早")
	} else {
		fmt.Println("两个日期相同")
	}

	// 注意事项：
	// 字符串格式必须固定： 日期格式必须严格遵守 "YYYY-MM-DD HH:mm:ss" 或类似的高到低顺序格式。
	//
	// 不支持复杂格式： 如果字符串中包含其他内容（如不同的时区、额外的符号），建议解析为 time.Time 再进行比较。
	//
	// 推荐场景： 如果数据源可靠、格式固定，直接用字符串比较是简单且高效的方式。但如果格式可能发生变化，使用 time.Parse 是更稳健的选择。
}
