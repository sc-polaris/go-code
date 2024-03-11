package main

import (
	"fmt"
	"os"
)

/*
Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，
我们通常用这个函数往文件中写入内容。

func Fprint(w io.Writer, a ...interface{}) (n int, err error)
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

0777表示：创建了一个普通文件，所有人拥有所有的读、写、执行权限

0666表示：创建了一个普通文件，所有人拥有对该文件的读、写权限，但是都不可执行

0644表示：创建了一个普通文件，文件所有者对该文件有读写权限，用户组和其他人只有读权限，
都没有执行权限

注意，只要满足io.Writer接口的类型都支持写入。
*/

func main() {
	// 向标准输入写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	file, err := os.OpenFile("./testFprint.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name := "沙河小王子"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(file, "往文件中写入信息:%s", name)
}
