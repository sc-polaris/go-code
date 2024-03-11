package main

import (
	"fmt"
	"strings"
)

func main() {
	var num string
	var k int
	fmt.Scan(&num, &k)

	var stack []byte
	for i := range num {
		//digit := num[i]
		fmt.Printf("%T\n", num[i])
		for k > 0 && len(stack) > 0 && num[i] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}

	stack = stack[:len(stack)-k]
	res := strings.TrimLeft(string(stack), "0")
	if res == "" {
		res = "0"
	}

	fmt.Println(res)
}
