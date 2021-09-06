package main

// 将请求参数存储起来 请求参数为k v 格式
// 用map存储k v

// 对map进行并发写需要进行加锁

// 声明一个map数据结构
type RequestBody map[string]interface{}

// 对map对象进行操作
func (r RequestBody) Set(key string, value interface{}) {

}

func main() {

}
