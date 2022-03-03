package main

import (
	"fmt"
	"reflect"
)

/*
在Go语言中，使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type），
程序通过类型对象可以访问任意值的类型信息。
func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}
*/

/*
在反射中关于类型还划分为两种：类型（Type）和种类（Kind）。因为在Go语言中我们可以使用
type关键字构造很多自定义类型，而种类（Kind）就是指底层的类型，但在反射中，当需要区分
指针、结构体等大品种的类型时，就会用到种类（Kind）。
*/

type myInt int64

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

/*
Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。
在reflect包中定义的Kind类型如下
*/

type Kind uint

const (
	Invalid       Kind = iota // 非法类型
	Bool                      // 布尔型
	Int                       // 有符号整型
	Int8                      // 有符号8位整型
	Int16                     // 有符号16位整型
	Int32                     // 有符号32位整型
	Int64                     // 有符号64位整型
	Uint                      // 无符号整型
	Uint8                     // 无符号8位整型
	Uint16                    // 无符号16位整型
	Uint32                    // 无符号32位整型
	Uint64                    // 无符号64位整型
	Uintptr                   // 指针
	Float32                   // 单精度浮点数
	Float64                   // 双精度浮点数
	Complex64                 // 64位复数类型
	Complex128                // 128位复数类型
	Array                     // 数组
	Chan                      // 通道
	Func                      // 函数
	Interface                 // 接口
	Map                       // 映射
	Ptr                       // 指针
	Slice                     // 切片
	String                    // 字符串
	Struct                    // 结构体
	UnsafePointer             // 底层指针
)

func main() {
	/*
		var a float32 = 3.14
		reflectType(a) // type:float32
		var b int64 = 100
		reflectType(b) // type:int64
	*/

	var a *float32 // 指针
	var b myInt    // 自定义类型
	var c rune     // 类型别名
	reflectType(a) // type: kind:ptr
	reflectType(b) // type:myInt kind:int64
	reflectType(c) // type:int32 kind:int32

	type person struct {
		name string
		age  int
	}

	type book struct{ title string }

	var d = person{
		name: "沙河小王子",
		age:  18,
	}

	var e = book{title: "《跟着小王子学Go语言》"}
	reflectType(d) // type:person kind:struct
	reflectType(e) // type:book kind:struct
}
