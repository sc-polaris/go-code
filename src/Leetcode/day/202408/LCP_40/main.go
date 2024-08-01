package main

import "slices"

/*
	力扣挑战赛」心算项目的挑战比赛中，要求选手从 N 张卡牌中选出 cnt 张卡牌，若这 cnt 张卡牌数字总和为偶数，则选手成绩「有效」且得分为 cnt 张卡牌数字总和。
	给定数组 cards 和 cnt，其中 cards[i] 表示第 i 张卡牌上的数字。请帮参赛选手计算最大的有效得分。若不存在获取有效得分的卡牌方案，则返回 0。
*/

/*
	从大到小排序，累加前 cnt 个数，记作 s。
	分类讨论：
	· 如果 s 是偶数，直接返回。
	· 如果 s 是奇数，那么可以：
		· 从前 cnt 个数中去掉一个最小的奇数，从后 n−cnt 个数中加进来一个最大的偶数，这样得分就变成偶数了。
		· 从前 cnt 个数中去掉一个最小的偶数，从后 n−cnt 个数中加进来一个最大的奇数，这样得分就变成偶数了。
		· 两种情况取最大值
*/

func maxmiumScore(cards []int, cnt int) int {
	slices.SortFunc(cards, func(a, b int) int { return b - a })
	s := 0
	for _, v := range cards[:cnt] {
		s += v
	}
	if s%2 == 0 {
		return s
	}

	replaceSum := func(x int) int {
		for _, v := range cards[cnt:] {
			if v%2 != x%2 { // 找到一个最大的奇偶性和 x 不同的数
				return s - x + v // 用 v 替换 x
			}
		}
		return 0
	}

	x := cards[cnt-1]               // 要么是最小的奇数，要么是最小的偶数
	ans := replaceSum(x)            // 替换 x
	for i := cnt - 2; i >= 0; i-- { // 前 cnt-1 个数
		if cards[i]%2 != x%2 { // 找到一个最小的奇偶性和 x 不同的数
			ans = max(ans, replaceSum(cards[i])) // 替换
			break
		}
	}
	return ans
}
