package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

// 遍历字符串
func traversalString() {
	s := "hello郑州"
	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for i, r := range s { // rune
		fmt.Printf("%d %v(%c) ", i, r, r)
	}
	fmt.Println()
}

// 修改字符串
func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}

// 类型转换
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	// 整型
	// 十进制
	var a int = 10
	fmt.Printf("%d\n", a) // 10
	fmt.Printf("%b\n", a) // 1010	占位符%b表示二进制

	// 八进制 以0开头
	var b int = 077
	fmt.Printf("%o\n", b) // 77

	// 十六进制 以0x开头
	var c int = 0xff
	fmt.Printf("%x\n", c) // ff
	fmt.Printf("%X\n", c) // FF

	// 浮点数
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	// 复数
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1, c2)

	// 布尔值

	// 字符串
	s1 := "hello"
	s2 := "你好"
	fmt.Println(s1, s2)

	// 转义符
	/*
		\r:回车符（返回行首）
		\n:换行符（直接跳到下一行的同列位置）
		\t:制表符
		\':单引号
		\":双引号
		\\:反斜杠
	*/
	// 打印文件路径
	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")
	// 多行字符串 使用反引号`
	s3 := `第一行
第二行
第三行`
	fmt.Println(s3)
	fmt.Println(len(s3))
	s3 += "123"
	s4 := "a b c d e f"
	strings := strings.Split(s4, " ")
	for _, s := range strings {
		fmt.Println(s)
	}

	// byte和rune类型
	// uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
	// rune类型，代表一个 UTF-8字符。
	var aa = '中'
	var bb = 'x'
	fmt.Println(aa, bb)
	traversalString()
	changeString()

	// 类型转换
	sqrtDemo()

	// 作业：编写代码统计出字符串"hello沙河小王子"中汉字的数量。
	ss2 := "hello沙河小王子"
	var count int
	for _, r := range ss2 {
		if unicode.Is(unicode.Han, r) {
			count++
		}
	}
	fmt.Printf("%T %d\n", ss2, count)
}
