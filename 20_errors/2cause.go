package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

var ErrNotFound = errors.New("err")
var ErrPermission = errors.New("ErrPermission")

func main() {
	err := ReadConfig() //
	err = parseArgs([]string{"1", "2"})
	if err != nil { // 在逻辑处理最顶层把堆栈信息直接打出来
		fmt.Printf("original error: %T %v \n", errors.Cause(err), errors.Cause(err)) // %T 打印错误类型 %v打印错误本身 与 sentinel error 进行 == 判断
		fmt.Printf("stack trace error:  %+v \n", err)                                // %+v 打印错误堆栈信息
	}
	//err = fmt.Errorf("access denied: %w", ErrPermission) // 加入额外信息
	//
	//if errors.Is(err, ErrNotFound) {
	//	// 将错误与sentinel 错误 进行比较
	//}
}

func ReadConfig() error {

	return nil
}

func parseArgs(args []string) error {
	if len(args) < 3 {
		return errors.New("not ....") // 遇到错误 使用errors.New 和 errors.Errorf 来返回错误 因为自己业务代码导致的错误
	}
	return nil
}

func WrapError(path string) error {
	_, err := os.Open(path)
	if err != nil {
		// 如果不能处理就直接往上抛  能处理就降级
		return errors.Wrapf(err, "failed to open %q", path) // 与第三方库进行协作时，使用wrapf包装错误   在自己打业务代码里面进行包装 基础库不建议
	}
	return nil
}
