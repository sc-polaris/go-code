package main

import "fmt"

/*
*printf系列函数都支持format格式化参数，在这里我们按照占位符将被替换的变量类型划分，方便查询和记忆。

占位符		说明
%v		值的默认格式表示
%+v		类似%v，但输出结构体时会添加字段名
%#v		值的Go语法表示
%T		打印值的类型
%%		百分号
*/

func main() {
	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	o := struct{ name string }{"小王子"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")
}
