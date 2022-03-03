package main

import (
	"fmt"
	"reflect"
)

// reflect.ValueOf()返回的是reflect.Value类型，其中包含了原始值的值信息。
// reflect.Value与原始值之间可以互相转换。

/*
reflect.Value类型提供的获取原始值的方法如下：
		方法										说明
Interface() interface {}		将值以 interface{} 类型返回，可以通过类型断言转换为指定类型
Int() 		int64				将值以 int 类型返回，所有有符号整型均可以此方式返回
Uint() 		uint64				将值以 uint 类型返回，所有无符号整型均可以此方式返回
Float() 	float64				将值以双精度（float64）类型返回，所有浮点数（float32、float64）均可以此方式返回
Bool() 		bool				将值以 bool 类型返回
Bytes() 	[]bytes				将值以字节数组 []bytes 类型返回
String() 	string				将值以字符串类型返回
*/

// 通过反射获取值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

// 通过反射设置变量的值
// 想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，
// 必须传递变量地址才能修改变量值。而反射中使用专有的Elem()方法来获取指针对应的值。

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) // 修改的副本，reflect包会引发panic
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

/*
func (v Value) IsNil() bool
IsNil()报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。
func (v Value) IsValid() bool
IsValid()返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。
*/

// 结构体反射
/*
任意值通过reflect.TypeOf()获得反射对象信息后，如果它的类型是结构体，
可以通过反射值对象（reflect.Type）的NumField()和Field()方法获得结构体成员的详细信息。

reflect.Type中与获取结构体成员相关的的方法如下表所示。
		方法																说明
Field(i int) StructField										根据索引，返回索引对应的结构体字段的信息。
NumField() int													返回结构体成员字段数量。
FieldByName(name string) (StructField, bool)					根据给定字符串返回字符串对应的结构体字段的信息。
FieldByIndex(index []int) StructField							多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息。
FieldByNameFunc(match func(string) bool) (StructField,bool)		根据传入的匹配函数匹配需要的字段。
NumMethod() int													返回该类型的方法集中方法的数目
Method(int) Method												返回该类型方法集中的第i个方法
MethodByName(string)(Method, bool)								根据方法名返回该类型方法集中的方法
*/

/*
// StructField类型
type StructField struct {
	// Name是字段的名字。PkgPath是非导出字段的包路径，对导出字段该字段为""。
	// 参见http://golang.org/ref/spec#Uniqueness_of_identifiers
	Name      string
	PkgPath   string
	Type      Type      // 字段的类型
	Tag       StructTag // 字段的标签
	Offset    uintptr   // 字段在结构体中的字节偏移量
	Index     []int     // 用于Type.FieldByIndex时的索引切片
	Anonymous bool      // 是否匿名字段
}
*/

// 结构体反射

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)

func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

// 遍历打印x包含的方法。
func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method type:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args []reflect.Value
		v.Method(i).Call(args)
	}
}

/*
反射是把双刃剑
反射是一个强大并富有表现力的工具，能让我们写出更灵活的代码。但是反射不应该被滥用，原因有以下三个。
	1. 基于反射的代码是极其脆弱的，反射中的类型错误会在真正运行的时候才会引发panic，
       那很可能是在代码写完的很长时间之后。
	2. 大量使用反射的代码通常难以理解。
	3. 反射的性能低下，基于反射实现的代码通常比正常代码运行速度慢一到两个数量级。
*/

func main() {
	/*
		var a float32 = 3.14
		var b int64 = 100
		reflectValue(a) // type is float32, value is 3.140000
		reflectValue(b) // type is int64, value is 100
		// 将int类型的原始值转换为reflect.Value类型
		c := reflect.ValueOf(10)
		fmt.Printf("type c :%T\n", c) // type c :reflect.Value
	*/

	/*
		var a int64 = 100
		//reflectSetValue1(a) //  reflect.Value.SetInt using unaddressable value
		reflectSetValue2(&a)
		fmt.Println(a)
	*/

	/*
		// IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效。
		// *int类型空指针
		var a *int
		fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
		// nil值
		fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
		// 实例化一个匿名结构体
		b := struct{}{}
		// 尝试从结构体中查找"abc"字段
		fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
		// 尝试从结构体中查找"abc"方法
		fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
		// map
		c := map[string]int{}
		// 尝试从map中查找一个不存在的键
		fmt.Println("map中不存在的键:", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())
	*/

	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n",
			field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n",
			scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}

	printMethod(stu1)

}
