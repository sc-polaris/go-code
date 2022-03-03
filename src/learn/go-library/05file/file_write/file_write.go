package main

import (
	"fmt"
	"os"
)

/*
os.OpenFile()函数能够以指定模式打开文件，从而实现文件写入相关功能。
func OpenFile(name string, flag int, perm FileMode) (*File, error) {
	...
}
其中：
	name：要打开的文件名 flag：打开文件的模式。 模式有以下几种
模式				含义
os.O_WRONLY		只写
os.O_CREATE		创建文件
os.O_RDONLY		只读
os.O_RDWR		读写
os.O_TRUNC		清空
os.O_APPEND		追加
perm：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01。
0777表示：创建了一个普通文件，所有人拥有所有的读、写、执行权限
0666表示：创建了一个普通文件，所有人拥有对该文件的读、写权限，但是都不可执行
0644表示：创建了一个普通文件，文件所有者对该文件有读写权限，用户组和其他人只有读权限，都没有执行权限
*/

func main() {
	file, err := os.OpenFile("../write.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "hello 沙河"
	file.Write([]byte(str))       // file.Write([]byte(str))
	file.WriteString("hello 小王子") // 直接写入字符串数据
}
