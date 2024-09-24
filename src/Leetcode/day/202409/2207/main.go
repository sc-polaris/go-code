package main

/*
	给你一个下标从 0 开始的字符串 text 和另一个下标从 0 开始且长度为 2 的字符串 pattern ，两者都只包含小写英文字母。

	你可以在 text 中任意位置插入 一个 字符，这个插入的字符必须是 pattern[0] 或者 pattern[1] 。注意，这个字符可以插入在 text 开头或者结尾的位置。

	请你返回插入一个字符后，text 中最多包含多少个等于 pattern 的 子序列 。

	子序列 指的是将一个字符串删除若干个字符后（也可以不删除），剩余字符保持原本顺序得到的字符串。
*/

/*
	贪心：
	首先计算插入之前的 pattern 子序列的个数，然后计算因为插入字母额外增加的 pattern 子序列的个数。

	设 x=pattern[0], y=pattern[1]。

	遍历 text 统计答案：遇到 y 时，如果左边出现了 3 个 x，那么就意味着我们找到了 3 个 pattern 子序列，把 3 加入答案。
	一般地，在遍历 text 的同时，维护 x 的出现次数 cntX。遇到 y 时，把 cntX 加入答案。

	然后考虑插入字母。

	根据题意，x 插入的位置越靠左，pattern 子序列的个数越多；y 插入的位置越靠右，pattern 子序列的个数越多。那么 x 应插在
	text 最左侧，y 应插在 text 最右侧。

	分类讨论：
	1. 把 x 插在 text 最左侧：答案额外增加 cntY，其中 cntY 是 y 在 text 中的出现次数。
	2. 把 y 插在 text 最右侧：答案额外增加 cntX，其中 cntX 是 x 在 text 中的出现次数。

	⚠注意：代码没有特判 x=y 的情况，要先更新答案，再更新 cntX，这可以保证更新答案时 cntX 表示的是当前字母左边的 x 的出现次数，cntX 尚未计入当前字母。
*/

func maximumSubsequenceCount(text string, pattern string) (ans int64) {
	x, y := pattern[0], pattern[1]
	cntX, cntY := 0, 0
	for i := range text {
		c := text[i]
		if c == y {
			ans += int64(cntX)
			cntY++
		}
		if c == x {
			cntX++
		}
	}
	return ans + int64(max(cntX, cntY))
}
