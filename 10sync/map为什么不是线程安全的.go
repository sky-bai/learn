package main

import "sync"

// 一般我们是家在配置文件 然后在启动程序

var icons map[string]string

// 初始化 全局map

func loadicons() {
	icons = map[string]string{
		"github": "image.Pt(64,64),",
		"google": "image.Pt(20, 20)",
	}
}

// Icon 我会先去判断全局的icons是否是已经初始化了的 然后再去访问icons里面的数据
func Icon(name string) string {
	if icons == nil {
		loadicons()
	}
	return icons[name]
}

// ---------------------------------------------------

// 程序在执行loadicons的时候 是分两步走的
func load() {
	// 1.先初始化map
	icons := make(map[string]string)
	// 2.再初始化图片
	icons["github"] = "image.Pt(64,64),"
	icons["google"] = "image.Pt(20, 20)"

	// 也就是说万一有一个goroutine 在初始化map 这时候全局map已经被初始化了 但是访问里面的数据可能就没有
	// 也就是说我需要将map的创建和 读写 这两个操作放在一起
}

// ==========================================================
var once sync.Once

// once 里面有个标志位 如果标志位为true 就不会再执行load方法了 如果false就表示是没有执行过 然后就会加锁

func once1(name string) string {
	once.Do(loadicons)
	return icons[name]
}

// do 方法将 map的创建和 访问放在一起

// 4200 - 500 = 3700
// 3700 - 3300 = 400

// partition 分区 文件夹 topic的内容存不下了 broker 一台服务器节点
