package main

import (
	"fmt"
	"runtime"
	"time"
)

/* https://www.cnblogs.com/sunsky303/p/9705727.html */

/*
可增长的栈：
	OS线程（操作系统线程）一般都有固定的栈内存（通常为2MB）,一个goroutine的栈
在其生命周期开始时只有很小的栈（典型情况下2KB），goroutine的栈不是固定的，他可
以按需增大和缩小，goroutine的栈大小限制可以达到1GB，虽然极少会用到这么大。所以
在Go语言中一次创建十万左右的goroutine也是可以的。

goroutine调度：
GPM是Go语言运行时（runtime）层面的实现，是go语言自己实现的一套调度系统。区别于操作系统调度OS线程。
	· G很好理解，就是个goroutine的，里面除了存放本goroutine信息外 还有与所在P的绑定等信息。
	  存储了goroutine的执行stack信息、goroutine状态以及goroutine的任务函数等；另外G对象是可以重用的。
	· P:表示逻辑processor管理着一组goroutine队列，P里面会存储当前goroutine运行的上下文环境（函数指针，堆栈地址及地址边界），
	  P会对自己管理的goroutine队列做一些调度（比如把占用CPU时间较长的goroutine暂停、运行后续的goroutine
	  等等）当自己的队列消费完了就去全局队列里取，如果全局队列里也消费完了会去其他P的队列里抢任务。
	· M（machine）是Go运行时（runtime）对操作系统内核线程的虚拟， M与内核线程一般是一一映射的关系，
	  一个goroutine最终是要放到M上执行的；
	  M代表着真正的执行计算资源。在绑定有效的p后，进入schedule循环；而schedule循环的机制大致是从各种队列、
	  p的本地队列中获取G，切换到G的执行栈上并执行G的函数，调用goexit做清理工作并回到m，如此反复。M并不保留
	  G状态，这是G可以跨M调度的基础。

P与M一般也是一一对应的。他们关系是： P管理着一组G挂载在M上运行。当一个G长久阻塞在一个M上时，
runtime会新建一个M，阻塞G所在的P会把其他的G 挂载在新建的M上。当旧的G阻塞完成或者认为其已经
死掉时回收旧的M。

P的个数是通过runtime.GOMAXPROCS设定（最大256），Go1.5版本之后默认为物理线程数。
在并发量大的时候会增加一些P和M，但不会太多，切换太频繁的话得不偿失。

单从线程调度讲，Go语言相比起其他语言的优势在于OS线程是由OS内核来调度的，goroutine
则是由Go运行时（runtime）自己的调度器调度的，这个调度器使用一个称为m:n调度的技术（
复用/调度m个goroutine到n个OS线程）。 其一大特点是goroutine的调度是在用户态下完成
的， 不涉及内核态与用户态之间的频繁切换，包括内存的分配与释放，都是在用户态维护着一块
大的内存池， 不直接调用系统的malloc函数（除非内存池需要改变），成本比调度OS线程低很
多。 另一方面充分利用了多核的硬件资源，近似的把若干goroutine均分在物理线程上， 再加
上本身goroutine的超轻量，以上种种保证了go调度方面的性能。
*/

/*
G:
	G是Goroutine的缩写，相当于操作系统中的进程控制块，在这里就是Goroutine的控制结构，
是对Goroutine的抽象。其中包括执行的函数指令及参数；G保存的任务对象；线程上下文切换，现
场保护和现场恢复需要的寄存器(SP、IP)等信息。
M:
	M是一个线程或称为Machine，所有M是有线程栈的。如果不对该线程栈提供内存的话，系统会
给该线程栈提供内存(不同操作系统提供的线程栈大小不同)。当指定了线程栈，则M.stack→G.stack，
M的PC寄存器指向G提供的函数，然后去执行。
P:
	P(Processor)是一个抽象的概念，并不是真正的物理CPU。所以当P有任务时需要创建或者唤
醒一个系统线程来执行它队列里的任务。所以P/M需要进行绑定，构成一个执行单元。
	P决定了同时可以并发任务的数量，可通过GOMAXPROCS限制同时执行用户级任务的操作系统线
程。可以通过runtime.GOMAXPROCS进行指定。在Go1.5之后GOMAXPROCS被默认设置可用的核数，
而之前则默认为1。
*/

/*
GOMAXPROCS：
	Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码。默
	认值是机器上的CPU核心数。例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS
	线程上（GOMAXPROCS是m:n调度中的n）。

	Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数。

	Go1.5版本之前，默认使用的是单核心执行。Go1.5版本之后，默认使用全部的CPU逻辑核心数。
*/

/* 我们可以通过将任务分配到不同的CPU逻辑核心上实现并行的效果，这里举个例子： */

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

/*
Go语言中的操作系统线程和goroutine的关系：
	1. 一个操作系统线程对应用户态多个goroutine。
	2. go程序可以同时使用多个操作系统线程。
	3. goroutine和OS线程是多对多的关系，即m:n。
*/

func main() {
	//runtime.GOMAXPROCS(1) // 两个任务只有一个逻辑核心，此时是做完一个任务再做另一个任务。
	runtime.GOMAXPROCS(2) // 将逻辑核心数设为2，此时两个任务并行执行
	go a()
	go b()
	time.Sleep(time.Second)
}
