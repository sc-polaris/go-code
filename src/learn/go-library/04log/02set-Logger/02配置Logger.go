package main

import (
	"fmt"
	"log"
	"os"
)

/*
· 标准logger的配置
	默认情况下的logger只会提供日志的时间信息，但是很多情况下我们希望得到更多信息，比如记录该日志的
文件名和行号等。log标准库中为我们提供了定制这些设置的方法。

log标准库中的
Flags函数会返回标准logger的输出配置，
SetFlags函数用来设置标准logger的输出配置。
	func Flags() int
	func SetFlags(flag int)

· flag选项
log标准库提供了如下的flag选项，它们是一系列定义好的常量。
const (
    // 控制输出日志信息的细节，不能控制输出的顺序和格式。
    // 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate         = 1 << iota     // 日期：2009/01/23
    Ltime                         // 时间：01:23:23
    Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
    Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
    Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
    LUTC                          // 使用UTC时间
    LstdFlags     = Ldate | Ltime // 标准logger的初始值
)

· 配置日志前缀：
log标准库中还提供了关于日志信息前缀的两个方法：
	func Prefix() string
	func SetPrefix(prefix string)
其中Prefix函数用来查看标准logger的输出前缀，SetPrefix函数用来设置输出前缀。

· 配置日志输出位置
	func SetOutput(w io.Writer)
SetOutput函数用来设置标准logger的输出目的地，默认是标准错误输出。
例如，下面的代码会把日志输出到同目录下的xx.log文件中。
func main() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[小王子]")
	log.Println("这是一条很普通的日志。")
}

如果你要使用标准的logger，我们通常会把上面的配置操作写到init函数中。
func init() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[小王子]")
}
*/

func init() {
	logFile, err := os.OpenFile("./log_demo.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[小王子]")
}

func main() {
	log.Println("这是一条很普通的日志。")
	log.Println("这是一条很普通的日志。")
}
