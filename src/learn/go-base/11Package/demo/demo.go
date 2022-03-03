package demo

import "fmt"

// 包级别标识符的可见性

// num 定义一个全局整型变量
// 首字母消协，对外不可见（只能在当前包内使用）
var num = 100

// Mode 定义一个常量
// 首字母大写，对外可见（可在其他包中使用）
const Mode = 1

// person 定义一个代表人的结构体
// 首字母小写，对外不可见（只能在当前包内使用）
type person struct {
	name string
	Age  int
}

// Add 返回两个整数和的函数
// 首字母大写，对外可见（可在其他包中使用）
func Add(x, y int) int {
	return x + y
}

func sayHi() {
	var myName = "梁壮" // 函数局部变量，只能在当前函数内使用
	fmt.Println(myName)
}

// 同样的规则也适用于结构体，结构体中可导出字段的字段名称必须首字母大写。

type Student struct {
	Name  string // 可在包外访问的字段
	class string // 仅限包内访问的字段
}
