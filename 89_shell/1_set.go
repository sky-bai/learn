package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		set -e: 表示此命令后，当某命令返回值非0时，将出错。如果是非交互环境，将直接退出，不再执行后续命令
		set +e: set -e的反向操作，恢复bash shell的默认行为，命令失败后继续执行后续命令

		set -u: (set -o nounset), 表示此命令之后，当某命令使用了未定义变量或参数时(特殊参数“@”和“*”除外)，将打印错误信息。如果是非交互环境(通常为脚本中)，将直接退出，不再执行后续命令
		set +u: set -u的反向操作，恢复bash shell的默认行为，命令使用未定义变量或参数时，继续执行后续命令
	*/

	fmt.Println(time.Now().Add(3600 * 24 * time.Hour * 3))
	// 当前时间加上3年
	fmt.Println(time.Now().AddDate(3, 0, 0))
}
