package main

import "strings"

/*
	给你一个用字符串表示的正整数 num ，请你以字符串形式返回不含尾随零的整数 num 。
*/

func removeTrailingZeros(num string) string {
	return strings.TrimRight(num, "0")
}
