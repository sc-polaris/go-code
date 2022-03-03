package main

import (
	"context"
	"fmt"
)

/*
context包中还定义了四个With系列函数。
· WithCancel
WithCancel的函数签名如下：
	func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
WithCancel返回带有新Done通道的父节点的副本。当调用返回的cancel函数或当关闭父上下文的Done通道时，
将关闭返回上下文的Done通道，无论先发生什么情况。
取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel。

上面的示例代码中，gen函数在单独的goroutine中生成整数并将它们发送到返回的通道。 gen的调用者在使用
生成的整数之后需要取消上下文，以免gen启动的内部goroutine发生泄漏。
*/

/*
context.Background() 返回一个空的Context
我们可以用这个 空的 Context 作为 goroutine 的root 节点（如果把整个 goroutine 的关系看作 树状）
使用context.WithCancel(parent)函数，创建一个可取消的子Context
函数返回值有两个：子Context Cancel 取消函数
例如：
	ctx, cancel := context.WithCancel(context.Background())
*/

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // return结束该goroutine，防止泄露
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 当我们取完需要的整数后调用cancel

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
