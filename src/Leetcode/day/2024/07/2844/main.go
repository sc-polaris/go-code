package main

import "strings"

/*
	给你一个下标从 0 开始的字符串 num ，表示一个非负整数。
	在一次操作中，您可以选择 num 的任意一位数字并将其删除。请注意，如果你删除 num 中的所有数字，则 num 变为 0。
	返回最少需要多少次操作可以使 num 变成特殊数字。
	如果整数 x 能被 25 整除，则该整数 x 被认为是特殊数字。

	一个数能被 25 整除，有如下五种情况：
	· 这个数是 0。
	· 这个数的末尾是 00，例如 100。
	· 这个数的末尾是 25，例如 225。
	· 这个数的末尾是 50，例如 350。
	· 这个数的末尾是 75，例如 475。
*/

/*
	方法一：枚举末尾
	首先，根据题目说的，我们可以把 num 中的所有数字都删除，得到 0，这需要删除 n 次。但如果 num 中有 0，那么删除 n−1 也可以得到 0。

	接下来，以示例 1 为例，看能否删除成末尾是 50 的数：
	1. 从右往左遍历 num，找最右边的 0。如果没有找到，或者最右边的 0 的下标是 0，则说明无法做到。
	2. 继续向左遍历，找最右边的 5，设其下标为 i。如果没有找到，则说明无法做到。
	3. 删除这个 5 右边的所有非 0 数字，这样就得到了一个以 50 结尾的字符串。
	4. 删除次数为 n−i−2。例如示例 1 中 5 的下标是 i=3，需要删除 7−3−2=2 次。
	其余 00,25,75 的计算方式同理，取 n−i−2 的最小值作为答案。
*/

func minimumOperations(num string) int {
	ans := len(num)
	if strings.Contains(num, "0") {
		ans-- // 可以删除 len(num)-1 次得到 "0"
	}
	f := func(tail string) {
		i := strings.LastIndexByte(num, tail[1])
		if i <= 0 {
			return
		}
		i = strings.LastIndexByte(num[:i], tail[0])
		if i < 0 {
			return
		}
		ans = min(ans, len(num)-i-2)
	}
	f("00")
	f("25")
	f("50")
	f("75")
	return ans
}

/*
	方法二：一次遍历
	在从右往左遍历的过程中：
	· 在之前找到 0 的情况下，如果当前数字 num[i] 是 0 或者 5，则立刻返回 n−i−2。
	· 在之前找到 5 的情况下，如果当前数字 num[i] 是 2 或者 7，则立刻返回 n−i−2。
	· 否则，如果 num[i] 是 0，标记我们找到了 0。
	· 否则，如果 num[i] 是 5，标记我们找到了 5。
	· 如果循环中没有返回，则最后返回 n 或者 n−1，取决于我们是否找到了 0。
*/

func minimumOperations2(num string) int {
	n := len(num)
	var fount0, fount5 bool
	for i := n - 1; i >= 0; i-- {
		c := num[i]
		if fount0 && (c == '0' || c == '5') ||
			fount5 && (c == '2' || c == '7') {
			return n - i - 2
		}
		if c == '0' {
			fount0 = true
		} else if c == '5' {
			fount5 = true
		}
	}

	if fount0 {
		return n - 1
	}
	return n
}
