package main

import (
	"fmt"
	"sync"
)

var wg1 sync.WaitGroup

/*
WaitGroup 对象内部有一个计数器，最初从0开始，它有三个方法：Add(), Done(), Wait() 用来控制计数器的数量。
Add(n) 把计数器设置为n ，Done() 每次把计数器-1 ，Wait() 会阻塞代码的运行，直到计数器地值减为0。
*/

func hello2(i int) {
	defer wg1.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg1.Add(1) // 启动一个goroutine就登记+1
		go hello2(i)
	}
	wg1.Wait() // 等待所有登记的goroutine都结束
	/*
		多次执行上面的代码，会发现每次打印的数字的顺序都不一致。
		这是因为10个goroutine是并发执行的，而goroutine的调度是随机的。
	*/
}
