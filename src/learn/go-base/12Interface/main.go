package main

import "fmt"

type ZhiFuBao struct {
	// 支付宝
}

type WeChat struct {
	// 微信
}

// Pay 支付宝的支付方法
func (z *ZhiFuBao) Pay(amount int64) {
	fmt.Printf("使用支付宝付款：%.2f元。\n", float64(amount/100))
}

// Pay 微信的支付方法
func (w *WeChat) Pay(amount int64) {
	fmt.Printf("使用微信付款：%.2f元。\n", float64(amount/100))
}

// CheckoutWithZFB 支付宝结账
func CheckoutWithZFB(obj *ZhiFuBao) {
	// 支付100元
	obj.Pay(100)
}

// CheckoutWithWX 微信支付结账
func CheckoutWithWX(obj *WeChat) {
	// 支付100元
	obj.Pay(100)
}

// Payer 包含支付方法的接口类型
type Payer interface {
	Pay(int64 int64)
}

func Checkout(obj Payer) {
	// 支付 100元
	obj.Pay(100)
}

// Sayer 接口
type Sayer interface {
	Say()
}

// Mover 接口
type Mover interface {
	Move()
}

// Cat 猫结构体类型
type Cat struct{}

// Dog 狗结构体类型
type Dog struct {
	Name string
}

// Say 实现Sayer借口
func (d Dog) Say() {
	fmt.Printf("%s会叫汪汪汪～\n", d.Name)
}

// Move 使用值接收者定义Move方法实现Mover接口
func (d Dog) Move() {
	fmt.Printf("%s会动\n", d.Name)
}

func (c Cat) Say() {
	fmt.Println("喵喵喵～")
}

// Move 使用指针接收者定义Move方法实现Mover接口
func (c *Cat) Move() {
	fmt.Println("猫会动")
}

// Car 汽车结构体类型
type Car struct {
	Brand string // 品牌
}

// Move Car类型实现Mover接口
func (c Car) Move() {
	fmt.Printf("%s速度70迈\n", c.Brand)
}

// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer // 嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}

// 空接口'

// Any 不包含任何方法的空接口类型
type Any interface{}

// Dog2 狗结构体
type Dog2 struct{}

// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	/*
		Checkout(&ZhiFuBao{}) // 之前调用支付宝支付

		Checkout(&WeChat{}) // 现在支持使用微信支付

		var x Sayer // 声明一个Sayer类型的变量x
		a := Cat{}  // 声明一个Cat类型变量a
		b := Dog{}  // 声明一个Dog类型变量b
		x = a       // 可以把Cat类型变量直接赋值给x
		x.Say()     // 喵喵喵
		x = b       // 可以把Dog类型变量直接赋值给x
		x.Say()     // 汪汪汪
	*/

	/*
		var x Mover // 声明一个Mover类型的变量x

		var d1 = Dog{} // d1是Dog类型
		x = d1         // 可以将d1赋值给变量x
		x.Move()

		var d2 = &Dog{} // d2是Dog指针类型
		x = d2          // 也可以将d2赋值给变量x
		x.Move()

		var c1 = &Cat{} // c1是*Cat类型
		x = c1          // 可以将c1当成Mover类型
		x.Move()
	*/

	/*
		// 下面的代码无法通过编译
		var c2 = Cat{} // c2是Cat类型
		x = c2         // 不能将c2当成Mover类型
	*/

	/*
		var d = Dog{Name: "旺财"}

		var s Sayer = d
		var m Mover = d

		s.Say()  // 对Sayer类型调用Say方法
		m.Move() // 对Mover类型调用Move方法
	*/

	/*
		var obj Mover

		obj = Dog{Name: "旺财"}
		obj.Move()

		obj = Car{Brand: "宝马"}
		obj.Move()
	*/

	/*
		var x Any

		x = "你好" // 字符串型
		fmt.Printf("type:%T value:%v\n", x, x)
		x = 100 // int型
		fmt.Printf("type:%T value:%v\n", x, x)
		x = true // 布尔型
		fmt.Printf("type:%T value:%v\n", x, x)
		x = Dog{} // 结构体类型
		fmt.Printf("type:%T value:%v\n", x, x)
	*/

	// 通常我们在使用空接口类型时不必使用type关键字声明，
	// 可以像下面的代码一样直接使用interface{}。
	//var x interface{} // 声明一个空接口类型变量x

	// 使用空接口实现可以保存任意值的字典。
	// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "热河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)
}
