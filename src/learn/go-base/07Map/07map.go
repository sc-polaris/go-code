package main

import "fmt"

func main() {
	/*
		scoreMap := make(map[string]int, 8)
		scoreMap["张三"] = 90
		scoreMap["小明"] = 100
		fmt.Println(scoreMap)
		fmt.Println(scoreMap["小明"])
		fmt.Printf("type of a:%T\n", scoreMap)
	*/
	/*
		userInfo := map[string]string{
			"username": "LiangZhuang",
			"password": "123456",
		}
		fmt.Println(userInfo)
	*/
	// 判断某个键是否存在
	//scoreMap := make(map[string]int, 8)
	//scoreMap["张三"] = 90
	//scoreMap["小明"] = 100
	//scoreMap["娜扎"] = 60
	// 如果key存在ok为true，v为对应的值，不存在ok为false，v为值类型的零值
	/*
		v, ok := scoreMap["张三"]
		if ok {
			fmt.Println(v)
		} else {
			fmt.Println("查无此人")
		}
	*/
	/*
		for key, value := range scoreMap {
			fmt.Println(key, value)
		}
		for key := range scoreMap {
			fmt.Println(key)
		}
	*/
	// delete
	/*
		delete(scoreMap, "小明")
		for key, value := range scoreMap {
			fmt.Println(key, value)
		}
	*/
	// 按照指定顺序遍历map
	/*
		rand.Seed(time.Now().UnixNano()) // 初始化随机数种子

		var scoreMap = make(map[string]int, 200)

		for i := 0; i < 100; i++ {
			key := fmt.Sprintf("stu%02d", i) // 生成stu开头的字符串
			value := rand.Intn(100)
			scoreMap[key] = value
		}
		// 去除map中的所有key存入切片keys
		var keys = make([]string, 0, 200)
		for key := range scoreMap {
			keys = append(keys, key)
		}
		// 对切片进行排序
		sort.Strings(keys)
		// 按照排序后对key遍历map
		for _, key := range keys {
			fmt.Println(key, scoreMap[key])
		}
	*/

	// 元素为map类型的切片
	/*
		var mapSlice = make([]map[string]string, 3)
		for index, value := range mapSlice {
			fmt.Printf("index:%d value:%v\n", index, value)
		}
		fmt.Println("after init")
		// 对切片中对map元素进行初始化
		mapSlice[0] = make(map[string]string, 10)
		mapSlice[0]["name"] = "小王子"
		mapSlice[0]["password"] = "123456"
		mapSlice[0]["address"] = "郑州轻工业大学"
		for index, value := range mapSlice {
			fmt.Printf("index:%d value:%v\n", index, value)
		}
	*/
	// 值为切片类型的map
	/*
		var sliceMap = make(map[string][]string, 3)
		fmt.Println(sliceMap)
		fmt.Println("after init")
		key := "中国"
		value, ok := sliceMap[key]
		if !ok {
			value = make([]string, 0, 2)
		}
		value = append(value, "北京", "上海")
		sliceMap[key] = value
		fmt.Println(sliceMap)
	*/

	/*	作业1
		s := "how do you do"
		strs := strings.Split(s, " ")
		var words = make(map[string]int)
		for _, str := range strs {
			words[str]++
		}
		for word, count := range words {
			fmt.Println(word, count)
		}
	*/

	// 作业2
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
