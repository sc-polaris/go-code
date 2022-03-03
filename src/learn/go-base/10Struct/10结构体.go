package main

import (
	"fmt"
)

/*
Go语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。
Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性。
*/

// NewInt 类型定义
type NewInt int

// MyInt 类型别名
//type MyInt = int

// 类型别名
type byte = uint8
type rune = int32

type person struct {
	name string
	city string
	age  int8
}

// 构造函数
// Go语言的结构体没有构造函数，我们可以自己实现。
// 因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，
// 所以该构造函数返回的是结构体指针类型。
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

type person1 struct {
	name, city string
	age        int8
}

type test1 struct {
	a int8
	b int8
	c int8
	d int8
}

type student struct {
	name string
	age  int
}

// Person 方法和接收者
// Person 结构体
type Person struct {
	name string
	age  int8
}

// NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

/*
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    函数体
}
*/

// Dream Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

// SetAge 设置p的年龄
// 使用指针接收者
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

// SetAge2 设置p的年龄
// 使用值接收者
/*
当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。
在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，
无法修改接收者变量本身。
*/
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}

// 什么时候应该使用指针类型接收者
/*
1. 需要修改接收者中的值
2. 接收者是拷贝代价比较大的大对象
3. 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
*/

func test2() {
	// 方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。
	p1 := NewPerson("小王子", 25)
	fmt.Println(p1.age) // 25
	p1.SetAge(30)
	fmt.Println(p1.age) // 30

	p2 := NewPerson("小王子", 25)
	fmt.Println(p2.age) // 25
	p2.SetAge2(30)      // (*p1).SetAge2(30)
	fmt.Println(p2.age) // 25
}

// MyInt 将int定义为自定义MyInt类型
type MyInt int

func (m MyInt) SayHello() {
	fmt.Println("Hello，我是一个int。")
}

// Person2 结构体的匿名字段
// Person2 结构体Person2类型
type Person2 struct {
	string
	int
}

// 嵌套结构体

type Address struct {
	Province string
	City     string
}

// User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address Address
}

// User2 嵌套匿名字段
type User2 struct {
	Name    string
	Gender  string
	Address // 匿名字段
}

func main() {
	/*
		var a NewInt
		var b MyInt

		fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
		fmt.Printf("type of b:%T\n", b) //type of b:int
	*/

	/*
		var p1 person
		p1.name = "热河娜扎"
		p1.city = "北京"
		p1.age = 18
		fmt.Printf("p1=%v\n", p1)
		fmt.Printf("p1=%#v\n", p1)

		var p2 = new(person)
		p2.name = "小王子"
		p2.age = 28
		p2.city = "上海"
		fmt.Printf("%T\n", p2)
		fmt.Printf("p2=%#v\n", p2)

		// 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
		// p3.name = "七米"其实在底层是(*p3).name = "七米"，这是Go语言帮我们实现的语法糖。
		p3 := &person{}
		fmt.Printf("%T\n", p3)
		fmt.Printf("p3=%#v\n", p3)
		p3.name = "梁壮"
		p3.age = 21
		p3.city = "郑州"
		fmt.Printf("p3=%#v\n", p3)

		// 没有初始化的结构体，其成员变量都是对应其类型的零值。
		var p4 person
		fmt.Printf("p4=%#v\n", p4) //p4=main.person{name:"", city:"", age:0}

		// 使用键值对初始化
		p5 := person{
			name: "小王子",
			city: "北京",
			age:  18,
		}
		fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"小王子", city:"北京", age:18}

		// 也可以对结构体指针进行键值对初始化
		p6 := &person{
			name: "小王子",
			city: "北京",
			age:  18,
		}
		fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"小王子", city:"北京", age:18}

		// 当某些字段没有初始值的时候，该字段可以不写。
		p7 := &person{
			city: "北京",
		}
		fmt.Printf("p7=%#v\n", p7) //p7=&main.person{name:"", city:"北京", age:0}

		// 使用值的列表初始化
		// 初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值：
		// 使用这种格式初始化时，需要注意：
	*/
	/*
		1. 必须初始化结构体的所有字段。
		2. 初始值的填充顺序必须与字段在结构体中的声明顺序一致。
		3. 该方式不能和键值初始化方式混用。
	*/

	/*
		p8 := &person{
			"沙河娜扎",
			"北京",
			28,
		}
		fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"沙河娜扎", city:"北京", age:28}

		// 结构体占用一块连续的内存。
		n := test1{
			1, 2, 3, 4,
		}
		fmt.Printf("n.a %p\n", &n.a)
		fmt.Printf("n.b %p\n", &n.b)
		fmt.Printf("n.c %p\n", &n.c)
		fmt.Printf("n.d %p\n", &n.d)

		// 空结构体是不占用空间的。
		var v struct{}
		fmt.Println(unsafe.Sizeof(v))

		// 面试题
		m := make(map[string]*student)
		stus := []student{
			{name: "小王子", age: 18},
			{name: "娜扎", age: 23},
			{name: "大王八", age: 9000},
		}

		for _, stu := range stus {
			m[stu.name] = &stu
		}

		for k, v := range m {
			fmt.Println(k, "=>", v)
		}

		p9 := newPerson("张三", "沙河", 90)
		fmt.Printf("%#v\n", p9)

		test2()
	*/

	var m1 MyInt
	m1.SayHello()
	m1 = 100
	fmt.Printf("%v  %T\n", m1, m1) //100  main.MyInt

	p1 := Person2{
		"小王子",
		18,
	}
	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.string, p1.int)

	user1 := User{
		Name:   "小王子",
		Gender: "男",
		Address: Address{
			Province: "河南",
			City:     "郑州",
		},
	}
	fmt.Printf("user1=%#v\n", user1)

	var user2 User2
	user2.Name = "小王子"
	user2.Gender = "男"
	user2.Address.Province = "河南" // 匿名字段默认使用类型名作为字段名
	user2.City = "焦作"             // 匿名字段可以省略
	fmt.Printf("user2=%#v\n", user2)

}
