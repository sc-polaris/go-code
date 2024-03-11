package main

import (
	"fmt"
)

// Address2 地址结构体
type Address2 struct {
	Province   string
	City       string
	CreateTime string
}

// Email 邮箱结构体
type Email struct {
	Account    string
	CreateTime string
}

// User3 用户结构体
type User3 struct {
	Name   string
	Gender string
	Address2
	Email
}

/* 结构体的继承 */

// Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

// Dog 狗
type Dog struct {
	Feet    int8
	*Animal // 通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪汪汪~\n", d.name)
}

// 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

// Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

// Class 班级
type Class struct {
	Title    string
	Students []*Student
}

// 结构体标签

// Student2 学生
type Student2 struct {
	ID     int    `json:"id"` // 通过制定tag实现json序列化该字段时的key
	Gender string // json序列化是默认使用字段名作为key
	name   string // 私有不能被json包访问
}

type Person3 struct {
	name   string
	age    int8
	dreams []string
}

// 因为slice和map这两种数据类型都包含了指向底层数据的指针，
// 因此我们在需要复制它们时要特别注意。

func (p *Person3) SetDreams(dreams []string) {
	//p.dreams = dreams // 错误
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

func main() {
	/*
		var user3 User3
		user3.Name = "热河娜扎"
		user3.Gender = "男"
		// user3.CreateTime = "2019" //ambiguous selector user3.CreateTime
		user3.Address2.CreateTime = "2000" //指定Address结构体中的CreateTime
		user3.Email.CreateTime = "2000"    //指定Email结构体中的CreateTime
	*/
	/*
		d1 := &Dog{
			Feet: 4,
			Animal: &Animal{ // 注意嵌套的是结构体指针
				name: "乐乐",
			},
		}
		d1.wang()
		d1.move()
	*/

	/*
		c := &Class{
			Title:    "101",
			Students: make([]*Student, 0, 200),
		}
		for i := 0; i < 10; i++ {
			stu := &Student{
				Name:   fmt.Sprintf("stu%02d", i),
				Gender: "男",
				ID:     i,
			}
			c.Students = append(c.Students, stu)
		}
		// JSON序列化：结构体-->JSON格式的字符串
		data, err := json.Marshal(c)
		if err != nil {
			fmt.Println("json marshal failed")
			return
		}
		fmt.Printf("json:%s\n", data)
		// JSON反序列化：JSON格式的字符串-->结构体
		str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
		c1 := &Class{}
		err = json.Unmarshal([]byte(str), c1)
		if err != nil {
			fmt.Println("json unmarshal failed!")
			return
		}
		fmt.Printf("%#v\n", c1)
	*/
	/*
		s1 := Student2{
			ID:     1,
			Gender: "男",
			name:   "热河娜扎",
		}
		data, err := json.Marshal(s1)
		if err != nil {
			fmt.Println("json marshal failed!")
			return
		}
		fmt.Printf("json str:%s\n", data) // json str:{"id":1,"Gender":"男"}
	*/

	p1 := Person3{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)

	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p1.dreams) // ?
}
