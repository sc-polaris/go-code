package main

/*
	用一个变量 less 维护目前子串中有 less 种字母的出现次数小于 t 中字母的出现次数。

	具体来说（注意下面算法中的 less 变量）：
	1. 初始化 ansLeft=−1, ansRight=m，用来记录最短子串的左右端点，其中 m 是 s 的长度。
	2. 用一个哈希表（或者数组）cntT 统计 t 中每个字母的出现次数。
	3. 初始化 left=0，以及一个空哈希表（或者数组）cntS，用来统计 s 子串中每个字母的出现次数。
	4. 初始化 less 为 t 中的不同字母个数。
	5. 遍历 s，设当前枚举的子串右端点为 right，把字母 c=s[right] 的出现次数加一。加一后，如果 cntS[c]=cntT[c]，
	   说明 c 的出现次数满足要求，把 less 减一。
	6. 如果 less=0，说明 cntS 中的每个字母及其出现次数都大于等于 cntT 中的字母出现次数，那么：
		a. 如果 right−left<ansRight−ansLeft，说明我们找到了更短的子串，更新 ansLeft=left, ansRight=right。
		b. 把字母 x=s[left] 的出现次数减一。减一前，如果 cntS[x]=cntT[x]，说明 x 的出现次数不满足要求，把 less 加一。
		c. 左端点右移，即 left 加一。
		d. 重复上述三步，直到 less>0，即 cntS 有字母的出现次数小于 cntT 中该字母的出现次数为止。
	7. 最后，如果 ansLeft<0，说明没有找到符合要求的子串，返回空字符串，否则返回下标 ansLeft 到下标 ansRight 之间的子串。
	代码实现时，可以把 cntS 和 cntT 合并成一个 cnt，定义
					cnt[x]=cntT[x]−cntS[x]
	如果 cnt[x]=0，就意味着窗口内字母 x 的出现次数和 t 的一样多。
*/

func minWindow(s string, t string) string {
	ansL, ansR := -1, len(s)
	var cnt [128]int
	less := 0
	for _, c := range t {
		if cnt[c] == 0 {
			less++ // 有 less 种字母的出现次数 < t 中的字母出现次数
		}
		cnt[c]++
	}

	l := 0
	for r, c := range s {
		cnt[c]--         // 右端字母移入子串
		if cnt[c] == 0 { // 原来窗口内 c 的出现次数比 t 的少，现在一样多
			less--
		}
		for less == 0 { // 涵盖：所有字母的出现次数都是 >=
			if r-l < ansR-ansL { // 找到更短的子串
				ansL, ansR = l, r
			}
			x := s[l]
			if cnt[x] == 0 {
				// x 移出窗口之前，检查出现次数，
				// 如果窗口内 x 的出现次数和 t 一样，
				// 那么 x 移出窗口后，窗口内 x 的出现次数比 t 的少
				less++
			}
			cnt[x]++ // 左端点字母移出子串
			l++
		}
	}
	if ansL < 0 {
		return ""
	}
	return s[ansL : ansR+1]
}
