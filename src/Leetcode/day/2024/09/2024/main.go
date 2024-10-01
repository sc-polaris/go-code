package main

/*
	一位老师正在出一场由 n 道判断题构成的考试，每道题的答案为 true （用 'T' 表示）或者 false （用 'F' 表示）。老师想增加
	学生对自己做出答案的不确定性，方法是 最大化 有 连续相同 结果的题数。（也就是连续出现 true 或者连续出现 false）。

	给你一个字符串 answerKey ，其中 answerKey[i] 是第 i 个问题的正确结果。除此以外，还给你一个整数 k ，表示你能进行以下
	操作的最多次数：
	· 每次操作中，将问题的正确答案改为 'T' 或者 'F' （也就是将 answerKey[i] 改为 'T' 或者 'F' ）。
	请你返回在不超过 k 次操作的情况下，最大 连续 'T' 或者 'F' 的数目。
*/

/*
	题意：
	求 answerKey 的一个最长子串，至多包含 k 个 T 或者至多包含 k 个 F。

	思路：
	由于子串越长，越无法满足要求，有单调性，可以用滑动窗口解决。
	1. 遍历 answerKey，枚举子串右端点 right，同时维护最小左端点 left 以及子串中的字符个数 cnt。
	2. 把 answerKey[right] 出现的次数加一。
	3. 如果 T 和 F 的出现次数都超过 k，那么必须不断移动左端点 left，同时减少 answerKey[left] 的出现次数，直到 T 和 F 的
	   出现次数至少有一个 ≤ k。
	4. 循环结束后，说明子串右端点在 right 时，对应的最小左端点为 left，用子串长度 right-left+1 更新答案的最大值。
	5. 遍历 answerKey 结束后，返回答案。

	代码实现时，由于 T 和 F 的 ASCII 值除以 2 后的奇偶性不同，也就是它们二进制的次低位不同，可以改为统计二进制次低位
*/

func maxConsecutiveAnswers(answerKey string, k int) (ans int) {
	cnt := [2]int{}
	l := 0
	for r, ch := range answerKey {
		cnt[ch>>1&1]++
		for cnt[0] > k && cnt[1] > k {
			cnt[answerKey[l]>>1&1]--
			l++
		}
		ans = max(ans, r-l+1)
	}
	return
}
