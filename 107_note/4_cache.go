package _07_note

import "container/list"

type LRU struct {
	// 使用双向链表实现
	List *list.List

	// 查找元素要高效使用map
	keys map[int]*list.Element

	// 容量
	Cap int
}

type pair struct {
	K, V int
}
