package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	str := "hello go"
	err := ioutil.WriteFile("../write.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
