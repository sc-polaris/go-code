package main

import "fmt"

/*
占位符			说明
%b		无小数部分、二进制指数的科学计数法，如-123456p-78
%e		科学计数法，如-1234.456e+78
%E		科学计数法，如-1234.456E+78
%f		有小数部分但无指数部分，如123.456
%F		等价于%f
%g		根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
%G		根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
*/

func main() {
	f := 12.34
	fmt.Printf("%b\n", f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%E\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%G\n", f)
}
