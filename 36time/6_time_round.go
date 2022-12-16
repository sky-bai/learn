package main

import (
	"fmt"
	"time"
)

func main() {
	duration := 125 * time.Millisecond
	rounded := duration.Round(time.Millisecond)
	fmt.Println(rounded) // prints "130ms"

	duration = 45 * time.Second
	rounded = duration.Round(time.Minute)
	fmt.Println(rounded) // prints "60s"
}

// 在这个示例中，我们将一个 125 毫秒的 time.Duration 值舍入到最接近的整毫秒数，并将一个 45 秒的 time.Duration 值舍入到最接近的整分钟数。

// time.Round() 函数可以在需要舍入时间值的任何场合使用。例如，您可能会在统计数据时使用它，以便将每小时的总计数据舍入到最接近的整小时数。您也可能会在处理计时数据时使用它，以便将每个任务的执行时间舍入到最接近的整分钟数。
//
//此外，还可以使用 time.Round() 函数来确保在计算时间差时，两个时间值的单位相同。例如，如果您需要计算两个时间值之间的时间差，您可能希望将其舍入到最接近的整小时数，以便更容易比较。
//
//总之，time.Round() 函数是一个非常有用的工具，可以用于许多不同的场合，帮助您将时间值舍入到最接近的整数倍的时间单位。
