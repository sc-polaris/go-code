package main

import (
	"errors"
	"fmt"
)

/*
Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。

func Errorf(format string, a ...interface{}) error
通常使用这种方式来自定义错误类型，例如：

err := fmt.Errorf("这是一个错误")
Go1.13版本为fmt.Errorf函数新加了一个%w占位符用来生成一个可以包裹Error的Wrapping Error。
*/

func main() {
	e := errors.New("原始错误error")
	w := fmt.Errorf("wrap了一个错误%w", e)
	fmt.Println(w)
}
