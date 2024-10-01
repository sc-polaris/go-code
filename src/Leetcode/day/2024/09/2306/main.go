package main

/*
	给你一个字符串数组 ideas 表示在公司命名过程中使用的名字列表。公司命名流程如下：

	1. 从 ideas 中选择 2 个 不同 名字，称为 ideaA 和 ideaB 。
	2. 交换 ideaA 和 ideaB 的首字母。
	3. 如果得到的两个新名字 都 不在 ideas 中，那么 ideaA ideaB（串联 ideaA 和 ideaB ，中间用一个空格分隔）是一个有效的公司名字。
	4. 否则，不是一个有效的名字。

	返回 不同 且有效的公司名字的数目。
*/

/*
	方法一：按照首字母分组
	按照首字母，把 ideas 分成（至多）26 组字符串。
	例如 ideas=[aa,ab,ac,bc,bd,be] 分成如下两组（只记录去掉首字母后的字符串）：
	· A 组（集合）：{a,b,c}。
	· B 组（集合）：{c,d,e}。

	B 组（集合）：{c,d,e}。
	1. 从 A 中选一个不等于 c 的字符串，这有 2 种选法。
	2. 从 B 中选一个不等于 c 的字符串，这有 2 种选法。
	3. 考虑两个字符串的先后顺序（谁在左谁在右），有 2 种方法。

	根据乘法原理，有 2×2×2=8 对符合要求的字符串。

	由于无法选交集中的字符串，一般地，从 A 和 B 中可以选出
				2⋅(∣A∣−∣A∩B∣)⋅(∣B∣−∣A∩B∣)

	对符合要求的字符串。其中 ∣S∣ 表示集合 S 的大小。
	枚举所有组对，计算上式，累加到答案中。
*/

func distinctNames(ideas []string) (ans int64) {
	group := [26]map[string]bool{}
	for i := range group {
		group[i] = map[string]bool{}
	}
	for _, s := range ideas {
		group[s[0]-'a'][s[1:]] = true // 按照首字母分组
	}

	for i, a := range group {
		for _, b := range group[:i] {
			m := 0 // 交集的大小
			for s := range a {
				if b[s] {
					m++
				}
			}
			ans += int64(len(a)-m) * int64(len(b)-m)
		}
	}
	return ans * 2
}
