package main

/*
	给你两个字符串 a 和 b，请返回 这两个字符串中 最长的特殊序列  的长度。如果不存在，则返回 -1 。

	「最长特殊序列」 定义如下：该序列为 某字符串独有的最长子序列（即不能是其他字符串的子序列） 。

	字符串 s 的子序列是在从 s 中删除任意数量的字符后可以获得的字符串。
	· 例如，"abc" 是 "aebdc" 的子序列，因为删除 "aebdc" 中斜体加粗的字符可以得到 "abc" 。 "aebdc" 的子序列还包括 "aebdc" 、 "aeb" 和 "" (空字符串)。
*/

/*
	分类讨论：
	· 如果 a = b，正如示例 3 所言，字符串 a 的每个子序列也是字符串 b 的每个子序列，字符串 b 的每个子序列
	  也是字符串 a 的子序列，不存在独有的子序列，返回 -1.
	· 如果 a != b：
		· 如果 a 比 b 长，答案的最大值为 a 的长度，能否做到？可以，把字符串 a 作为独有子序列，它不是 b 的子序列，所以答案为 a 的长度。
		· 如果 b 比 a 长，同上，答案为 b 的长度。
		· 如果 a 和 b 一样长，由于 a != b，把字符串 a 作为独有序列，它不是 b 的子序列，所以答案为 a 的长度。
		· 所以，当 a != b 时，返回 a 和 b 二者的最大长度。
*/

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	return max(len(a), len(b))
}
