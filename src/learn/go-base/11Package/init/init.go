package main

import "fmt"

// 每一个包的初始化是先从初始化包级别变量开始的。
// 例如从下面的示例中我们就可以看出包级别变量的初始化会先于init初始化函数。

var x int8 = 10

const pi = 3.14

func init() {
	fmt.Println("x:", x)
	fmt.Println("pi:", pi)
	sayHi()
}

func sayHi() {
	fmt.Println("Hello world!")
}

func main() {
	fmt.Println("你好，世界！")
}
