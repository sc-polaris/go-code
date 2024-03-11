package main

import "fmt"

/*
占位符		说明
%p		表示为十六进制，并加上前导的0x
*/

func main() {
	a := 10
	fmt.Printf("%p\n", &a)
	fmt.Printf("%#p\n", &a)
}
