package main

import "fmt"

func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}

func main() {
	/* 初始化数组时可以使用初始化列表来设置数组元素的值。 */
	/*var testArray [3]int
	var numArray = [3]int{1, 2}
	var cityArray = [3]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)
	fmt.Println(numArray)
	fmt.Println(cityArray)*/
	/* 2.*/
	/*var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)
	fmt.Println(numArray)
	fmt.Printf("type of numArray:%T\n", numArray)   //type of numArray:[2]int
	fmt.Println(cityArray)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray) //type of cityArray:[3]strin

	a := [...]int{1: 1, 3, 5}
	fmt.Println(a)
	fmt.Printf("type of a:%T\n", a)*/

	/*var a = [...]string{"北京", "上海", "深圳"}
	// 方法1：for循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	// 方法2：range遍历
	for index, value := range a {
		fmt.Println(index, value)
	}*/

	/*a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a)
	fmt.Println(a[2][1])

	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}*/

	/*a := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}*/
	a := [3]int{10, 20, 30}
	modifyArray(a) //在modify中修改的是a的副本x
	fmt.Println(a) //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b) //在modify中修改的是b的副本x
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]
}
