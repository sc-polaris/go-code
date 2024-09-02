package main

/*
	给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

	注意：
	· 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
	· 如果 s 中存在这样的子串，我们保证它是唯一的答案。

*/

/*
	方法一：O(52m+n)
	我们枚举 s 子串的右端点 right（子串最后一个字母的下标），如果子串涵盖 t，就不断移动左端点 left 直到不涵盖为止。在移动
	过程中更新最短子串的左右端点。

	具体来说：
	1. 初始化 ansLeft = -1, ansRight := m，用来记录最短子串的左右端点，其中 m 是 s 的长度。
	2. 用一个哈希表（或者数组） cntT 统计 t 中每个字母的出现次数。
	3. 初始化 left = 0，以及一个空哈希表（或者数组）cntS，用来统计 s 子串中每个字母的出现次数。
	4. 遍历 s，设当前枚举的子串右端点为 right，把 s[right] 的出现次数加一。
	5. 遍历 cntS 中的每个字母及其出现次数，如果出现次数都大于等于 cntT 中的字母出现次数：
		a. 如果 right-left < ansRight-ansLeft，说明我们找到了更短的子串，更新 ansLeft = left,ansRight = right。
		b. 把 s[left] 出现的次数减一。
		c. 左端点右移，即 left 加一。
		d. 重复上述三步，直到 cntS 有字母的出现次数小于 cntT 中该字母的出现次数为止。
	6. 最后，如果 ansLeft < 0，说明没有找到符合要求的子串，返回空字符串，否则返回下标 ansLeft 到下标 ansRight 之间的子串。

	由于本题大写字母和小写字母都有，为了方便，代码实现时可以直接创建大小为 128 的数组，保证所有 ASCII 字符都可以统计。
*/

func minWindow(s string, t string) string {
	isCovered := func(cntS, cntT []int) bool {
		for i := 'A'; i <= 'Z'; i++ {
			if cntS[i] < cntT[i] {
				return false
			}
		}
		for i := 'a'; i <= 'z'; i++ {
			if cntS[i] < cntT[i] {
				return false
			}
		}
		return true
	}

	ansLeft, ansRight, left := -1, len(s), 0
	var cntS, cntT [128]int
	for _, c := range t {
		cntT[c]++
	}
	for right, c := range s { // 移动子串右端点
		cntS[c]++                         // 右端点字母移入子串
		for isCovered(cntS[:], cntT[:]) { // 涵盖
			if right-left < ansRight-ansLeft { // 找到更短的子串
				ansLeft, ansRight = left, right // 记录此时的左右端点
			}
			cntS[s[left]]-- // 左端点字母移出子串
			left++          // 移动子串左端点
		}
	}
	if ansLeft < 0 {
		return ""
	}
	return s[ansLeft : ansRight+1]
}

/*
	方法二：优化
	上面的代码每次都要花费 O(∣Σ∣) 的时间去判断是否涵盖，能不能优化到 O(1) 呢？

	可以。用一个变量 less 维护目前子串中有 less 种字母的出现次数小于 t 中字母的出现次数。
	具体来说（注意下面算法中的 less 变量）：
	1. 初始化 ansLeft=−1, ansRight=m，用来记录最短子串的左右端点，其中 m 是 s 的长度。
	2. 用一个哈希表（或者数组）cntT 统计 t 中每个字母的出现次数。
	3. 初始化 left=0，以及一个空哈希表（或者数组）cntS，用来统计 s 子串中每个字母的出现次数。
	4. 初始化 less 为 t 中的不同字母个数。
	5. 遍历 s，设当前枚举的子串右端点为 right，把字母 c=s[right] 的出现次数加一。加一后，如果 cntS[c]=cntT[c]，说明
	   c 的出现次数满足要求，把 less 减一。
	6. 如果 less=0，说明 cntS 中的每个字母及其出现次数都大于等于 cntT 中的字母出现次数，那么：
		a. 如果 right−left<ansRight−ansLeft，说明我们找到了更短的子串，更新 ansLeft=left, ansRight=right。
		b. 把字母 x=s[left] 的出现次数减一。减一前，如果 cntS[x]=cntT[x]，说明 x 的出现次数不满足要求，把 less 加一。
		c. 左端点右移，即 left 加一。
		d. 重复上述三步，直到 less>0，即 cntS 有字母的出现次数小于 cntT 中该字母的出现次数为止。
	7. 最后，如果 ansLeft<0，说明没有找到符合要求的子串，返回空字符串，否则返回下标 ansLeft 到下标 ansRight 之间的子串。
*/

func minWindow2(s, t string) string {
	ansLeft, ansRight := -1, len(s)
	left, less := 0, 0
	var cntS, cntT [128]int
	for _, c := range t {
		if cntT[c] == 0 {
			less++ // 有 less 种字母的出现次数 < t 中的字母出现次数
		}
		cntT[c]++
	}
	for right, c := range s { // 移动子串右端点
		cntS[c]++
		if cntS[c] == cntT[c] {
			less-- // c 的出现次数从 < 变成 >=
		}
		for less == 0 { // 涵盖：所有字母的出现次数都是 >=
			if right-left < ansRight-ansLeft { // 找到更短的子串
				ansLeft, ansRight = left, right // 记录此时的左右端点
			}
			x := s[left] // 左端点字母
			if cntS[x] == cntT[x] {
				less++ // x 的出现次数从 >= 变成 <
			}
			cntS[x]-- // 左端点字母移出子串
			left++
		}
	}
	if ansLeft < 0 {
		return ""
	}
	return s[ansLeft : ansRight+1]
}
