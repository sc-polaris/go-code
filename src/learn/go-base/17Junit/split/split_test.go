package split

import (
	"fmt"
	"reflect"
	"testing"
)

/*
func TestSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := Split("a:b:c", ":")         // 程序输出的结果
	want := []string{"a", "b", "c"}    // 期望的结果
	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("expected:%v, got:%v\n", want, got) // 测试失败输出错误提示
	}
}
*/

/*
测试组
	我们现在还想要测试一下split函数对中文字符串的支持，这个时候我们可以再编写一个
TestChineseSplit测试函数，但是我们也可以使用如下更友好的一种方式来添加更多的测试用例。

	我们的测试出现了问题，仔细看打印的测试失败提示信息：
expected:[河有 又有河], got:[ 河有 又有河]，你会发现[ 河有 又有河]中有个不明显的空串，
这种情况下十分推荐使用%#v的格式化方式。
func TestSplit(t *testing.T) {
	// 定义一个测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}
	// 定义一个存储测试用例的切片
	tests := []test{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		{input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}
	// 遍历切片，逐一执行测试用例
	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			//t.Errorf("expected:%v, got:%v", tc.want, got)
			t.Errorf("expected:%#v, got:%#v", tc.want, got)
		}
	}
}
*/

/*
	子测试
	看起来都挺不错的，但是如果测试用例比较多的时候，
	我们是没办法一眼看出来具体是哪个测试用例失败了。我们可能会想到下面的解决办法
func TestSplit(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}
	for name, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("name:%s expected:%#v, got:%#v", name, tc.want, got) // 将测试用例的name格式化输出
		}
	}
}
	上面的做法是能够解决问题的。同时Go1.7+中新增了子测试，我们可以按照如下方式使用t.Run执行子测试：：

	我们都知道可以通过-run=RegExp来指定运行的测试用例，还可以通过/来指定要运行的子测试用例，例如：
	go test -v -run=Split/simple只会运行simple对应的子测试用例。
*/

/*
测试覆盖率：
Go提供内置功能来检查你的代码覆盖率。我们可以使用go test -cover来查看测试覆盖率。

Go还提供了一个额外的-coverprofile参数，用来将覆盖率相关的记录信息输出到一个文件。
例如：go test -cover -coverprofile=c.out

上面的命令会将覆盖率相关的信息输出到当前文件夹下面的c.out文件中，然后我们执行
go tool cover -html=c.out，使用cover工具来处理生成的记录信息，该命令会打开本地
的浏览器窗口生成一个HTML报告。

每个用绿色标记的语句块表示被覆盖了，而红色的表示没有被覆盖。
*/

/*
func TestSplit(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name:%s expected:%#v, got:%#v", name, tc.want, got) // 将测试用例的name格式化输出
			}
		})
	}
}
*/

/*
func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}
*/

/*
· 基准测试：
基准测试就是在一定的工作负载之下检测程序性能的一种方法。基准测试的基本格式如下：
func BenchmarkName(b *testing.B){
    // ...
}

基准测试以Benchmark为前缀，需要一个*testing.B类型的参数b，基准测试必须要执行b.N次，
这样的测试才有对照性，b.N的值是系统根据实际情况去调整的，从而保证测试的稳定性。testing.B拥有的方法如下：
unc (c *B) Error(args ...interface{})
func (c *B) Errorf(format string, args ...interface{})
func (c *B) Fail()
func (c *B) FailNow()
func (c *B) Failed() bool
func (c *B) Fatal(args ...interface{})
func (c *B) Fatalf(format string, args ...interface{})
func (c *B) Log(args ...interface{})
func (c *B) Logf(format string, args ...interface{})
func (c *B) Name() string
func (b *B) ReportAllocs()
func (b *B) ResetTimer()
func (b *B) Run(name string, f func(b *B)) bool
func (b *B) RunParallel(body func(*PB))
func (b *B) SetBytes(n int64)
func (b *B) SetParallelism(p int)
func (c *B) Skip(args ...interface{})
func (c *B) SkipNow()
func (c *B) Skipf(format string, args ...interface{})
func (c *B) Skipped() bool
func (b *B) StartTimer()
func (b *B) StopTimer()

· 基准测试示例
我们为split包中的Split函数编写基准测试如下：
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}
基准测试并不会默认执行，需要增加-bench参数，所以我们通过执行go test -bench=Split命令执行基准测试。

其中BenchmarkSplit-8表示对Split函数进行基准测试，数字8表示GOMAXPROCS的值，这个对于并发基准测
试很重要。10000000和203ns/op表示每次调用Split函数耗时203ns，这个结果是10000000次调用的平均值。

我们还可以为基准测试添加-benchmem参数，来获得内存分配的统计数据。
其中，112 B/op表示每次操作内存分配了112字节，3 allocs/op则表示每次操作进行了3次内存分配。

我们将我们的Split函数优化如下：
func Split(s, sep string) (result []string) {
	result = make([]string, 0, strings.Count(s, sep)+1)
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
这一次我们提前使用make函数将result初始化为一个容量足够大的切片，而不再像之前一样通过调用append
函数来追加。我们来看一下这个改进会带来多大的性能提升：
这个使用make函数提前分配内存的改动，减少了2/3的内存分配次数，并且减少了一半的内存分配。
*/

/*
· 重置时间：
b.ResetTimer之前的处理不会放到执行时间里，也不会输出到报告中，所以可以在之前做一些不计划
作为测试报告的操作。
func BenchmarkSplit(b *testing.B) {
	time.Sleep(5 * time.Second) // 假设需要做一些耗时的无关操作
	b.ResetTimer()              // 重置计时器
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}
*/

//func BenchmarkSplit(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Split("沙河有沙又有河", "沙")
//	}
//}

/*
· 并行测试：
func (b *B) RunParallel(body func(*PB))会以并行的方式执行给定的基准测试。

RunParallel会创建出多个goroutine，并将b.N分配给这些goroutine执行， 其中goroutine数量的默认
值为GOMAXPROCS。用户如果想要增加非CPU受限（non-CPU-bound）基准测试的并行性， 那么可以在
RunParallel之前调用SetParallelism 。RunParallel通常会与-cpu标志一同使用。

还可以通过在测试命令后添加-cpu参数如go test -bench=. -cpu 1来指定使用的CPU数量。
*/

//func BenchmarkSplitParallel(b *testing.B) {
//	// b.SetParallelism(1) // 设置使用的CPU数
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			Split("沙河有沙又有河", "沙")
//		}
//	})
//}

/*
· Setup与TearDown：
	测试程序有时需要在测试之前进行额外的设置（setup）或在测试之后进行拆卸（teardown）。
·· TestMain：
	通过在*_test.go文件中定义TestMain函数来可以在测试之前进行额外的设置（setup）或在测试之后
进行拆卸（teardown）操作。
	如果测试文件包含函数:func TestMain(m *testing.M)那么生成的测试会先调用 TestMain(m)，
然后再运行具体测试。TestMain运行在主goroutine中, 可以在调用 m.Run前后做任何设置（setup）和
拆卸（teardown）。退出测试的时候应该使用m.Run的返回值作为参数调用os.Exit。

	一个使用TestMain来设置Setup和TearDown的示例如下：
func TestMain(m *testing.M) {
	fmt.Println("write setup code here...") // 测试之前的做一些设置
	// 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
	retCode := m.Run()                         // 执行测试
	fmt.Println("write teardown code here...") // 测试之后做一些拆卸工作
	os.Exit(retCode)                           // 退出测试
}
	需要注意的是：在调用TestMain时, flag.Parse并没有被调用。所以如果TestMain 依赖于command-line
标志 (包括 testing 包的标记), 则应该显示的调用flag.Parse。

·· 子测试的Setup与Teardown
	有时候我们可能需要为每个测试集设置Setup与Teardown，也有可能需要为每个子测试设置Setup与Teardown。
下面我们定义两个函数工具函数如下：
// 测试集的Setup与Teardown
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:测试之后的teardown")
	}
}

// 子测试的Setup与Teardown
func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:子测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:子测试之后的teardown")
	}
}
*/

//func TestMain(m *testing.M) {
//	fmt.Println("write setup code here...") // 测试之前的做一些设置
//	// 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
//	retCode := m.Run()                         // 执行测试
//	fmt.Println("write teardown code here...") // 测试之后做一些拆卸工作
//	os.Exit(retCode)                           // 退出测试
//}

// 测试集的Setup与Teardown
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:测试之后的teardown")
	}
}

// 子测试的Setup与Teardown
func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:子测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:子测试之后的teardown")
	}
}

func TestSplit(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}
	teardownTestCase := setupTestCase(t) // 测试之前执行setup操作
	defer teardownTestCase(t)            // 测试之后执行testdoen操作

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			teardownSubTest := setupSubTest(t) // 子测试之前执行setup操作
			defer teardownSubTest(t)           // 测试之后执行testdoen操作
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

/*
· 示例函数：
	被go test特殊对待的第三种函数就是示例函数，它们的函数名以Example为前缀。它们既没有参数也没
有返回值。标准格式如下：
func ExampleName() {
    // ...
}

下面的代码是我们为Split函数编写的一个示例函数：
func ExampleSplit() {
	fmt.Println(split.Split("a:b:c", ":"))
	fmt.Println(split.Split("沙河有沙又有河", "沙"))
	// Output:
	// [a b c]
	// [ 河有 又有河]
}

为你的代码编写示例代码有如下三个用处：
	1. 示例函数能够作为文档直接使用，例如基于web的godoc中能把示例函数与对应的函数或包相关联。
	2. 示例函数只要包含了// Output:也是可以通过go test运行的可执行测试。
	3. 示例函数提供了可以直接运行的示例代码，可以直接在golang.org的godoc文档服务器上使用
       Go Playground运行示例代码。下图为strings.ToUpper函数在Playground的示例函数效果。
*/

func ExampleSplit() {
	fmt.Println()
	fmt.Println(Split("沙河有沙又有河", "沙"))
	//Output:
	//[a b c]
	//[ 河有 又有河]
}
