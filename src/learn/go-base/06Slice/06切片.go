package main

import (
	"fmt"
	"sort"
)

/*
切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。

切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。
*/

func main() {
	// 声明切片类型
	/*var a []string              // 声明一个字符串切片
	var b = []int{}             // 声明一个整型切片并初始化
	var c = []bool{false, true} // 声明一个布尔切片并初四话
	//var d = []bool{false, true} // 声明一个布尔切片并初四话
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a == nil)
	fmt.Println(b == nil)
	fmt.Println(c == nil)*/
	//fmt.Println(c == d) //切片是引用类型，不支持直接比较，只能和nil比较
	/*a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))*/
	/*
		a[2:]  // 等同于 a[2:len(a)]
		a[:3]  // 等同于 a[0:3]
		a[:]   // 等同于 a[0:len(a)]
	*/
	//s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)
	//fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
	// 完整切片
	//t := a[1:3:5]
	//fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))

	//var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
	//s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	//s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil

	//s1 := make([]int, 3) //[0 0 0]
	//s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	//s2[0] = 100
	//fmt.Println(s1) //[100 0 0]
	//fmt.Println(s2) //[100 0 0]

	// 切片遍历和数组是一致的
	/*
		s := []int{1, 3, 5}

		for i := 0; i < len(s); i++ {
			fmt.Println(i, s[i])
		}

		for index, value := range s {
			fmt.Println(index, value)
		}
	*/

	// append()方法为切片添加元素
	/*
		var s []int
		s = append(s, 1)       // [1]
		s = append(s, 2, 3, 4) // [1 2 3 4]
		s2 := []int{5, 6, 7}
		s = append(s, s2...) // [1 2 3 4 5 6 7]
		fmt.Println(s)
	*/
	//append()添加元素和切片扩容
	//var numSlice []int
	//for i := 0; i < 10; i++ {
	//	numSlice = append(numSlice, i)
	//	fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	//}

	// copy()复制切片
	/*
		a := []int{1, 2, 3, 4, 5}
		c := make([]int, 5, 5)
		copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
		fmt.Println(a) //[1 2 3 4 5]
		fmt.Println(c) //[1 2 3 4 5]
		c[0] = 1000
		fmt.Println(a) //[1 2 3 4 5]
		fmt.Println(c) //[1000 2 3 4 5]
	*/

	// 从切片中删除元素
	// Go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素。
	/*
		a := []int{30, 31, 32, 33, 34, 35, 36, 37}
		// 要删除索引为2的元素
		a = append(a[:2], a[3:]...)
		fmt.Println(a) //[30 31 33 34 35 36 37]
	*/
	// 练习题
	// 1.请写出下面代码的输出结果。
	/*
		var a = make([]string, 5, 10)
		for i := 0; i < 10; i++ {
			a = append(a, fmt.Sprintf("%v", i))
		}
		fmt.Println(a)
	*/

	// 内置排序
	a := [...]int{3, 7, 8, 9, 1}
	// 升序
	sort.Sort(sort.IntSlice(a[:]))
	fmt.Println(a)
	// 降序
	sort.Sort(sort.Reverse(sort.IntSlice(a[:])))
	fmt.Println(a)
}
