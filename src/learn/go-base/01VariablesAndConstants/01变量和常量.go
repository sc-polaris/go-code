package main

import "fmt"

// 全局变量
var m = 100

const (
	pi = 3.141592653589793
	e1 = 2.71828
)

const (
	n1 = iota // 常量计数器 0
	n2        // 1
	n3        // 2
	n4        // 3
)

const (
	m1 = iota // 0
	m2 = 100  // 100
	m3 = iota // 2
	m4        // 3
)

const m5 = iota // 0

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

const (
	a, b = iota + 1, iota + 2 // 1, 2
	c, d                      // 2, 3
	e, f                      // 3, 4
)

func foo() (int, string) {
	return 10, "Liangzhuang"
}

func main() {
	x, _ := foo()
	_, y := foo()
	fmt.Println(x, y)
	n := 100
	m := 200
	fmt.Println(m, n)
}
