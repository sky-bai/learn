package main

import (
	"fmt"
	"strings"
)

func main() {
	strllll := "中心领导|开发处室领导|测试处室领导|开发人员|测试人员"
	st2 := "中心领导|开发处室领导|"
	sr := strings.Split("中心领导|开发处室领导|测试处室领导|开发人员|测试人员", "|")
	fmt.Println("sr:", sr)
	is1 := strings.Contains(strllll, "中心领导")
	fmt.Println("is1:", is1)

	if strings.Contains(strllll, "开发处室领导") || strings.Contains(strllll, "测试处室领导") {
		fmt.Println("该角色是部门管理员")
	}
	if strings.Contains(st2, "开发处室领导") || strings.Contains(st2, "测试处室领导") {
		fmt.Println("该角色是部门管理员")
	}

}

func is1(s string) bool {
	return strings.Contains(s, "1")
}
