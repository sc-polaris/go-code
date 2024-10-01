package main

/*
	给定字符串列表 strs ，返回其中 最长的特殊序列 的长度。如果最长特殊序列不存在，返回 -1 。

	特殊序列 定义如下：该序列为某字符串 独有的子序列（即不能是其他字符串的子序列）。

	 s 的 子序列可以通过删去字符串 s 中的某些字符实现。

	例如，"abc" 是 "aebdc" 的子序列，因为您可以删除"aebdc"中的下划线字符来得到 "abc" 。"aebdc"的子序列还包括"aebdc"、 "aeb" 和 "" (空字符串)。
*/

/*
	我们需要枚举每个字符串的所有子序列吗？
	不需要，如果短的子序列是「独有子序列」，那么更长的子序列也会是「独有子序列」。或者说，子序列越长，越不可能是其他字符串的子序列。
	所以只需要枚举字符串 s = str[i]，判断 s 是否是其他字符串的子序列，如果不是，则用 s 的长度更新答案的最大值。
*/

// 判断 s 是否为 t 的子序列
func isSubSeq(s, t string) bool {
	i := 0
	for _, c := range t {
		if s[i] == byte(c) {
			i++
			if i == len(s) { // 所有字符串匹配完毕
				return true // s 是 t 的子序列
			}
		}
	}
	return false
}

func findLUSlength(strs []string) int {
	ans := -1
next:
	for i, s := range strs {
		if len(s) <= ans { // 不会让 ans 变大
			continue
		}
		for j, t := range strs {
			if j != i && isSubSeq(s, t) {
				continue next
			}
		}
		ans = len(strs[i])
	}
	return ans
}
