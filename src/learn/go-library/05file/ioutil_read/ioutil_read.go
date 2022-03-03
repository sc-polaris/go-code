package main

import (
	"fmt"
	"io/ioutil"
)

// io/ioutil包的ReadFile方法能够读取完整的文件，只需要将文件名作为参数传入。

// ioutil.ReadFile读取整个文件
func main() {
	content, err := ioutil.ReadFile("../read.txt")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}
