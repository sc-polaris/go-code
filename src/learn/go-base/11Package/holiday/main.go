// holiday/file_read.go

package main

import (
	"fmt"
	"github.com/lz397548123/hello/v2"
	//"github.com/q1mi/hello" // 导入当前项目下的包
	"holiday/summer" // 导入github上第三方包
)

func main() {
	fmt.Println("现在是假期时间...")
	//hello.SayHi()
	summer.Driving()
	hello.SayHi("张三")
}
