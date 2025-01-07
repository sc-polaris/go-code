package main

import "strings"

/*
	给你一个下标从 0 开始的字符串 s ，该字符串由用户输入。按键变更的定义是：使用与上次使用的按键不同的键。例如 s = "ab" 表示按键变更一次，而 s = "bBBb" 不存在按键变更。

	返回用户输入过程中按键变更的次数。

	注意：shift 或 caps lock 等修饰键不计入按键变更，也就是说，如果用户先输入字母 'a' 然后输入字母 'A' ，不算作按键变更。
*/

func countKeyChanges(s string) (ans int) {
	s = strings.ToLower(s)
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			ans++
		}
	}
	return
}

// 对于同一字母的大写和小写，ASCII 值的二进制的低 5 位是相同的，所以只需统计 s[i−1]&31 != s[i]&31

func countKeyChanges2(s string) (ans int) {
	for i := 1; i < len(s); i++ {
		if s[i-1]&31 != s[i]&31 {
			ans++
		}
	}
	return
}
