package main

import "fmt"

// 区别于C/C++中的指针，Go语言中的指针不能进行偏移和运算，是安全指针。

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

/*
new和make的区别
1. 二者都是用来做内存分配的。
2. make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
3. 而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
*/

func main() {
	/*
		a := 10
		modify1(a)
		fmt.Println(a)
		modify2(&a)
		fmt.Println(a)
	*/

	var a *int
	a = new(int)
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b = make(map[string]int, 10)
	b["沙河娜扎"] = 100
	fmt.Println(b)
}
