package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

/*
编写代码利用反射实现一个ini文件的解析器程序。
*/

type Student struct {
	Name string `ini:"name"`
	Age  int    `ini:"age"`
}

// ini配置读取到map中
func iniToMap(path string) (map[string]string, error) {
	var m = map[string]string{}

	file, err := os.Open(path)
	if err != nil {
		return m, err
	}
	defer file.Close()

	br := bufio.NewReader(file)

	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		var kv = strings.Split(string(line), ":")

		if len(kv) >= 2 {
			m[kv[0]] = kv[1]
		}
	}

	return m, nil
}

// 解析map到结构体
func parseConf(i interface{}, m map[string]string) {
	k := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	for s := 0; s < v.Elem().NumField(); s++ {
		if tagValue, ok := m[k.Elem().Field(s).Tag.Get("ini")]; ok {
			switch k.Elem().Field(s).Type.Kind() {
			case reflect.Int:
				num, _ := strconv.Atoi(tagValue)
				v.Elem().Field(s).SetInt(int64(num))
			case reflect.String:
				v.Elem().Field(s).SetString(tagValue)
			}
		}
	}

}

func main() {
	var m, err = iniToMap("./a.ini")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var student = Student{}
	parseConf(&student, m)
	fmt.Println(student)
}
