package main

import "fmt"

type Mover interface {
	Move()
}

type Dog struct {
	Name string
}

func (d *Dog) Move() {
	fmt.Println("狗在跑～")
}

type Car struct {
	Brand string
}

func (c *Car) Move() {
	fmt.Println("汽车在跑～")
}

// 如果对一个接口值有多个实际类型需要判断，推荐使用switch语句来实现。

// justifyType 对传入的空接口类型变量x进行类型断言
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string, value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type!")
	}
}

func main() {
	var m Mover
	fmt.Println(m == nil) // true
	// 我们不能对一个空接口值调用任何方法，否则会产生panic。->宕机
	//m.Move() // panic: runtime error: invalid memory address or nil pointer dereference

	m = &Dog{Name: "旺财"}
	m = new(Car)
	fmt.Println(m == nil) // false

	// 接口值是支持相互比较的，当且仅当接口值的动态类型和动态值都相等时才相等。
	var (
		x Mover = new(Dog)
		y Mover = new(Car)
	)
	fmt.Println(x == y) // false

	// 但是有一种特殊情况需要特别注意，如果接口值的保存的动态类型相同，
	// 但是这个动态类型不支持互相比较（比如切片），那么对它们相互比较时就会引发panic。
	//var z interface{} = []int{1, 2, 3}
	//fmt.Println(z == z) // panic: runtime error: comparing uncomparable type []int

	m = &Dog{Name: "旺财"}
	fmt.Printf("%T\n", m) // *main.Dog

	m = new(Car)
	fmt.Printf("%T\n", m) // *main.Car

	// 而想要从接口值中获取到对应的实际值需要使用类型断言，其语法格式如下。
	// x.(T) x:表示接口类型的变量 T:表示断言x可能是的类型
	// 该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。
	var n Mover = &Dog{Name: "旺财"}
	v, ok := n.(*Dog)
	if ok {
		fmt.Println("类型断言成功")
		v.Name = "富贵" // 变量v是*Dog类型
	} else {
		fmt.Println("类型断言失败")
	}

}
