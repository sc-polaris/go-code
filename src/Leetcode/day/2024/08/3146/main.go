package main

/*
	给你两个字符串 s 和 t，每个字符串中的字符都不重复，且 t 是 s 的一个排列。

	排列差 定义为 s 和 t 中每个字符在两个字符串中位置的绝对差值之和。

	返回 s 和 t 之间的 排列差 。
*/

func findPermutationDifference(s string, t string) (ans int) {
	pos := [26]int{}
	for i, c := range s {
		pos[c-'a'] = i
	}
	for i, c := range t {
		ans += abs(i - pos[c-'a'])
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
