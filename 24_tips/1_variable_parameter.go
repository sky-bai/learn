package _4_tips

// 使用可变参数去代替切片

// 这只是我编的一个例子，但它与我所写的很多代码相同。 这里的问题是他们假设他们会被调用于多个条目。 但是很多时候这些类型的函数只用一个参数调用，为了满足函数参数的要求，它必须打包到一个切片内。
//
//另外，因为 ids 参数是切片，所以你可以将一个空切片或 nil 传递给该函数，编译也没什么错误。 但是这会增加额外的测试负载，因为你应该涵盖这些情况在测试中。
//
//举一个这类 API 的例子，最近我重构了一条逻辑，要求我设置一些额外的字段，如果一组参数中至少有一个非零。 逻辑看起来像这样：

func anyPositive(values ...int) bool {
	for _, v := range values {
		if v > 0 {
			return true
		}
	}
	return false
}
func main() {
	//if svc.MaxConnections > 0 || svc.MaxPendingRequests > 0 || svc.MaxRequests > 0 || svc.MaxRetries > 0 {
	//}
	//
	//if anyPositive(svc.MaxConnections, svc.MaxPendingRequests, svc.MaxRequests, svc.MaxRetries) {
	//}
}

//func Save(f *os.File, doc *Document) error
//我可以指定这个函数 Save，它将 *os.File 作为写入 Document 的目标。但这样做会有一些问题
//
//Save 的签名排除了将数据写入网络位置的选项。假设网络存储可能在以后成为需求，则此功能的签名必须改变，从而影响其所有调用者。
//
//Save 测试起来也很麻烦，因为它直接操作磁盘上的文件。因此，为了验证其操作，测试时必须在写入文件后再读取该文件的内容。
//
//而且我必须确保 f 被写入临时位置并且随后要将其删除。
//
//*os.File 还定义了许多与 Save 无关的方法，比如读取目录并检查路径是否是符号链接。 如果 Save 函数的签名只用 *os.File 的相关内容，那将会很有用。
//
//我们能做什么 ？
// 使用 io.ReadWriteCloser，我们可以应用接口隔离原则来重新定义 Save 以获取更通用文件形式。 可读可写的文件是一个更好的选择，因为它允许我们将数据写入任何实现了 io.ReadWriteCloser 接口的目标，而不仅仅是 *os.File。
//
//通过此更改，任何实现 io.ReadWriteCloser 接口的类型都可以替换以前的 *os.File。
//
//这使 Save 在其应用程序中更广泛，并向 Save 的调用者阐明 *os.File 类型的哪些方法与其操作有关。
//
//而且，Save 的作者也不可以在 *os.File 上调用那些不相关的方法，因为它隐藏在 io.ReadWriteCloser 接口后面。
//
//但我们可以进一步采用接口隔离原则。
//
//首先，如果 Save 遵循单一功能原则，它不可能读取它刚刚写入的文件来验证其内容 - 这应该是另一段代码的功能。
//
//// Save writes the contents of doc to the supplied
//// WriteCloser.
//func Save(wc io.WriteCloser, doc *Document) error
//因此，我们可以将我们传递给 Save 的接口的规范缩小到只写和关闭。
//
//其次，通过向 Save 提供一个关闭其流的机制，使其看起来仍然像一个文件，这就提出了在什么情况下关闭 wc 的问题。
//
//可能 Save 会无条件地调用 Close，或者在成功的情况下调用 Close。
//
//这给 Save 的调用者带来了问题，因为它可能希望在写入文档后将其他数据写入流。
//
// Save writes the contents of doc to the supplied
// Writer.
//func Save(w io.Writer, doc *Document) error
