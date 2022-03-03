package main

import (
	"errors"
	"fmt"
	"strings"
)

func intSum(x, y int) int {
	return x + y
}

func sayHello() {
	fmt.Println("Hello 沙河")
}

func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

func intSum3(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}

func calc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

/*
func someFunc(x string) []int {
	if x == "" {
		return nil // 没必要返回[]{}
	}
}
*/

// 定义全局变量num
var num int64 = 10

func testGlobalVar() {
	fmt.Printf("num=%d\n", num)
}

func testLocalVar() {
	// 定义一个局部变量x，仅在该函数内生效
	var x int64 = 100
	fmt.Printf("x=%d\n", x)
}

func testNum() {
	num := 100
	fmt.Printf("num=%d\n", num) // 函数中优先使用局部变量
}

func testLocalVar2(x, y int) {
	fmt.Println(x, y) // 函数的参数也是只在本函数中生效
	if x > 0 {
		z := 100 // 变量z只在if语句块生效
		fmt.Println(z)
	}
	// fmt.Println(z)//此处无法使用变量z
}

func testLocalVar3() {
	for i := 0; i < 10; i++ {
		fmt.Println(i) // 变量i只在当前for语句块中生效
	}
	// fmt.Println(i) //此处无法使用变量i
}

// 定义函数类型
type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func calc2(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别对操作符")
		return nil, err
	}
}

// 闭包指的是一个函数和与其相关的引用环境组合而成的实体
// 闭包=函数+引用环境
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

// suffix:后缀
func makeSuffixFunc(suffix string) func(string string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc3(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func calc4(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// Go语言中目前（Go1.12）是没有异常机制，但是使用panic/recover模式来处理错误。
// panic可以在任何地方引发，但recover只有在defer调用的函数中有效。

func funcA() {
	fmt.Println("func A")
}

// recover()必须搭配defer使用。
// defer一定要在可能引发panic的语句之前定义。
func funcB() {
	//panic("panic in B")
	defer func() {
		err := recover()
		// 如果程序出现了panic错误，可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}

// 练习题 分金币
/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	for _, user := range users {
		for _, ch := range user {
			if ch == 'e' || ch == 'E' {
				distribution[user]++
				coins--
			} else if ch == 'i' || ch == 'I' {
				distribution[user] += 2
				coins -= 2
			} else if ch == 'o' || ch == 'O' {
				distribution[user] += 3
				coins -= 3
			} else if ch == 'u' || ch == 'U' {
				distribution[user] += 4
				coins -= 4
			}
		}
	}

	return coins
}

func main() {
	/*
		sayHello()
		ret := intSum(10, 20)
		fmt.Println(ret)
	*/
	/*
		ret1 := intSum2()
		ret2 := intSum2(10)
		ret3 := intSum2(10, 20)
		ret4 := intSum2(10, 20, 30)
		fmt.Println(ret1, ret2, ret3, ret4) //0 10 30 60
	*/

	/*
		ret5 := intSum3(100)
		ret6 := intSum3(100, 10)
		ret7 := intSum3(100, 10, 20)
		ret8 := intSum3(100, 10, 20, 30)
		fmt.Println(ret5, ret6, ret7, ret8) //100 110 130 160
	*/
	//testGlobalVar()
	//testLocalVar()
	//testNum()

	// add和sub都能赋值给calculation类型的变量。
	/*
		var c calculation               // 声明一个calculation类型的变量c
		c = add                         // 把add赋值给c
		fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
		fmt.Println(c(1, 2))            // 像调用add一样调用c

		f := add                        // 将函数add赋值给变量f1
		fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
		fmt.Println(f(10, 20))          // 像调用add一样调用f
	*/

	/*
		ret2 := calc2(10, 20, add)
		fmt.Println(ret2)
	*/

	/*
		// 将匿名函数保存到变量
		add := func(x, y int) {
			fmt.Println(x + y)
		}
		add(10, 20) // 通过变量调用匿名函数

		// 自执行函数：匿名函数定义完加()直接执行
		func(x, y int) {
			fmt.Println(x + y)
		}(10, 20)

		var f = adder()
		fmt.Println(f(10)) // 10
		fmt.Println(f(20)) // 30
		fmt.Println(f(30)) // 60

		f1 := adder()
		fmt.Println(f1(40)) // 40
		fmt.Println(f1(50)) // 90
	*/

	/*
		var f = adder2(10)
		fmt.Println(f(10)) // 20
		fmt.Println(f(20)) // 40

		f1 := adder2(20)
		fmt.Println(f1(40)) // 60
		fmt.Println(f1(50)) // 110

		jpgFunc := makeSuffixFunc(".jpg")
		txtFunc := makeSuffixFunc(".txt")
		fmt.Println(jpgFunc("test"))
		fmt.Println(txtFunc("test"))
	*/

	/*
		f1, f2 := calc3(10)
		fmt.Println(f1(1), f2(2)) // 11 9
		fmt.Println(f1(3), f2(4)) // 12 8
		fmt.Println(f1(5), f2(6)) // 13 7
	*/
	/*
		fmt.Println("start")
		defer fmt.Println(1)
		defer fmt.Println(2)
		defer fmt.Println(3)
		fmt.Println("end")

		A 1 2 3
		B 10 2 12
		BB 10 12 22
		AA 1 3 4
	*/

	/*
		fmt.Println(f1())
		fmt.Println(f2())
		fmt.Println(f3())
		fmt.Println(f4())
	*/

	/*
		x := 1
		y := 2
		defer calc4("AA", x, calc4("A", x, y))
		x = 10
		defer calc4("BB", x, calc4("B", x, y))
		y = 20
	*/

	/*
		funcA()
		funcB()
		funcC()
	*/

	left := dispatchCoin()
	fmt.Println("剩下：", left)

}
