package main

import "fmt"

// Go 语言中使用一个名为 error 接口来表示错误类型。
//type error interface {
//	Error() string
//}

/*
// 创造错误
func queryById(id int64) (*Info, error) {
	if id <= 0 {
		return nil, errors.New("无效的id")
	}
	// ...
}
*/

// 错误结构体类型

// OpError 自定义结构体类型
type OpError struct {
	Op string
}

func (e *OpError) Error() string {
	return fmt.Sprintf("无权执行%s操作", e.Op)
}

func main() {
	// 定义一个错误变量，例如标准库io.EOF错误定义如下。
	//var EOF = errors.New("EOF")

	// 当我们需要传入格式化的错误描述信息时，使用fmt.Errorf是个更好的选择。
	//fmt.Errorf("查询数据库失败，err:%v", err)
	/*
		但是上面的方式会丢失原有的错误类型，只拿到错误描述的文本信息。
		为了不丢失函数调用的错误链，使用fmt.Errorf时搭配使用特殊的格式化动词%w，
		可以实现基于已有的错误再包装得到一个新的错误。
		fmt.Errorf("查询数据库失败，err:%w", err)
	*/

	/*
		对于这种二次包装的错误，errors包中提供了以下三个方法。
		func Unwrap(err error) error                 // 获得err包含下一层错误
		func Is(err, target error) bool              // 判断err是否包含target
		func As(err error, target interface{}) bool  // 判断err是否为target类型
	*/

}
