package main

// 请求参数
type BodyMap map[string]interface{}

// 这里是在定义的时候 定义入参需要传一个函数 这个函数的入参为bodymap结构体
func (bm BodyMap) Set(key string, value func(bm BodyMap)) {

}
