package main

type User struct {
	ID   int64
	Name string
}

func main() {
	str := new(string)
	*str = "hello"

	_ = GetUserInfo()
}

func GetUserInfo() *User {
	return &User{ID: 1, Name: "张三"}
}

// 很关键的一点就是它有没有被作用域之外所引用，即作用域仍保留在main中，因此它没有逃逸。由此可见，是否被作用域之外的区域所引用是逃逸的重要原因之一。
// interface类型会导致该对象被分配到堆上
