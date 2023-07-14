// Command click is a chromedp example demonstrating how to use a selector to
// click on an element.
package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://pkg.go.dev/time`),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(`body > footer`),
		// find and click "Example" link
		chromedp.Click(`#example-After`, chromedp.NodeVisible),
		// retrieve the text of the textarea
		chromedp.Value(`#example-After textarea`, &example),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example:\n%s", example)
}

// `body > footer` 是一个 CSS 选择器，用于选择页面中 `<footer>` 元素的直接子元素。
//
//在 CSS 中，选择器是用于选择 HTML 元素的模式。其中，`>` 符号表示选择符前面的元素是选择符后面元素的直接子元素。
//
//解释 `body > footer` 的含义如下：
//
//- `body`: 选择 `<body>` 元素，表示选择整个页面的主体部分。
//- `>`: 选择符号，表示后面的元素是前面元素的直接子元素。
//- `footer`: 选择 `<footer>` 元素，表示选择 `<body>` 元素的直接子元素中的 `<footer>` 元素。
//
//因此，`body > footer` 表示选择页面中作为 `<body>` 元素直接子元素的 `<footer>` 元素。

// `#example-After` 是一个 CSS 选择器，用于选择具有特定 ID 属性的元素。
//
//在 CSS 中，ID 选择器用于选择具有特定 ID 值的元素。ID 是 HTML 元素的唯一标识符，每个 ID 在文档中应该是唯一的。
//
//解释 `#example-After` 的含义如下：
//
//- `#`: 选择符号，表示后面的值是一个 ID。
//- `example-After`: 具体的 ID 值。
//
//因此，`#example-After` 表示选择具有 ID 为 "example-After" 的元素。在示例代码中，`chromedp.Click("#example-After", chromedp.NodeVisible)` 使用该选择器选择页面上具有 ID "example-After" 的元素，并执行点击操作。

// 浮动元素是指在页面中具有浮动属性的元素。浮动元素会脱离正常的文档流，而是浮动在页面中的某个位置。
// 之前的样式只由本身元素支撑
