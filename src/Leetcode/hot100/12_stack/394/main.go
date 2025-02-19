package main

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	var numStack []int
	var strStack []string
	num := 0
	result := ""
	for _, c := range s {
		if c >= '0' && c <= '9' {
			n, _ := strconv.Atoi(string(c))
			num = num*10 + n
		} else if c == '[' {
			strStack = append(strStack, result)
			result = ""
			numStack = append(numStack, num)
			num = 0
		} else if c == ']' {
			count := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			str := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			result = str + strings.Repeat(result, count)
		} else {
			result += string(c)
		}
	}
	return result
}
