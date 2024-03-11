package main

import "fmt"

/*
Sprint系列函数会把传入的数据生成并返回一个字符串。

func Sprint(a ...interface{}) string
func Sprintf(format string, a ...interface{}) string
func Sprintln(a ...interface{}) string
*/

func main() {
	s1 := fmt.Sprint("沙河小娜扎")
	name := "沙河小娜扎"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("沙河小娜扎")
	fmt.Println(s1, s2, s3)
}
