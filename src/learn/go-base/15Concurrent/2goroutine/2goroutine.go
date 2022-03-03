package main

import (
	"fmt"
	"time"
)

/*
并发：同一时间段内执行多个任务（你在用微信和两个女朋友聊天）。
并行：同一时刻执行多个任务（你和你朋友都在用微信和女朋友聊天）。

	Go语言的并发通过goroutine实现。goroutine类似于线程，属于用户态的线程，
我们可以根据需要创建成千上万个goroutine并发工作。goroutine是由Go语言的
运行时（runtime）调度完成，而线程是由操作系统调度完成
	Go语言还提供channel在多个goroutine间进行通信。goroutine和channel是
Go 语言秉承的 CSP（Communicating Sequential Process）并发模式的重要实现基础。

	在java/c++中我们要实现并发编程的时候，我们通常需要自己维护一个线程池，并
且需要自己去包装一个又一个的任务，同时需要自己去调度线程执行任务并维护上下文切换，
这一切通常会耗费程序员大量的心智。那么能不能有一种机制，程序员只需要定义很多个任务，
让系统去帮助我们把这些任务分配到CPU上实现并发执行呢？

	Go语言中的goroutine就是这样一种机制，goroutine的概念类似于线程，但 goroutine
是由Go的运行时（runtime）调度和管理的。Go程序会智能地将 goroutine 中的任务合理地分
配给每个CPU。Go语言之所以被称为现代化的编程语言，就是因为它在语言层面已经内置了调度和
上下文切换的机制。

	在Go语言编程中你不需要去自己写进程、线程、协程，你的技能包里只有一个技能–goroutine，
当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，开启一个goroutine
去执行这个函数就可以了，就是这么简单粗暴。
*/

func hello() {
	fmt.Println("Hello Goroutine!")
}
func main() {
	/*
		这一次的执行结果只打印了main goroutine done!，并没有打印Hello Goroutine!。为什么呢？

		在程序启动时，Go程序就会为main()函数创建一个默认的goroutine。

		当main()函数返回的时候该goroutine就结束了，所有在main()函数中启动的goroutine会一同结束，
		main函数所在的goroutine就像是权利的游戏中的夜王，其他的goroutine都是异鬼，夜王一死它转化
		的那些异鬼也就全部GG了。

		所以我们要想办法让main函数等一等hello函数，最简单粗暴的方式就是time.Sleep了。
	*/
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	time.Sleep(time.Second)

	/*
		执行上面的代码你会发现，这一次先打印main goroutine done!，然后紧接着打印Hello Goroutine!。

		首先为什么会先打印main goroutine done!是因为我们在创建新的goroutine的时候需要花费一些时间，
		而此时main函数所在的goroutine是继续执行的。
	*/
}
