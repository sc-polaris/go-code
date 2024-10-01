package main

import (
	"slices"
)

/*
	给你一个长度为 n 的二维整数数组 items 和一个整数 k 。

	items[i] = [profiti, categoryi]，其中 profiti 和 categoryi 分别表示第 i 个项目的利润和类别。

	现定义 items 的 子序列 的 优雅度 可以用 total_profit + distinct_categories^2 计算，其中 total_profit 是子序列中所有项目的利润总和，distinct_categories
	是所选子序列所含的所有类别中不同类别的数量。

	你的任务是从 items 所有长度为 k 的子序列中，找出 最大优雅度 。

	用整数形式表示并返回 items 中所有长度恰好为 k 的子序列的最大优雅度。

	注意：数组的子序列是经由原数组删除一些元素（可能不删除）而产生的新数组，且删除不改变其余元素相对顺序。
*/

/*
	按照利润从大到小排序。先把前 k 个项目选上
	考虑第 k+1 个项目，如果要选它，我们必须从前 k 个项目中移除一个项目。
	由于已经按照利润从大到小排序，选这个项目不会让 totalProfit 变大，所以重点考虑能否让 distinctCategories 变大

	分类讨论：
	· 如果第 k+1 个项目和前面某个已选项目的类别相同，那么无论怎么移除都不会让 distinctCategories 变大，所以无需选择这个项目。
	· 如果第 k+1 个项目和前面任何已选项目的类别都不一样，考虑移除前面已选项目中的哪一个：
		· 如果移除的项目的类别只出现一次，那么选第 k+1 个项目后，distinctCategories 一减一增，保持不变，所以不考虑这种情况。
		· 如果移除的项目的类别重复出现多次，那么选第 k+1 个项目后，distinctCategories 会增加一，此时有可能会让优雅度变大，
	 	  一定要选择这个项目。为什么说「一定」呢？因为 totalProfit 只会变小，我们现在的目标就是让 totalProfit 保持尽量大，
		  同时让 distinctCategories 增加，那么只能让 distinctCategories 增加就立刻选上！因为后面的利润更小，现在不选的
		  话将来 totalProfit 只会更小。

	按照这个过程，继续考虑选择后面的项目。计算优雅度，取最大值，即为答案。

	代码实现时，我们应当移除已选项目中类别和前面重复且利润最小的项目，这可以用一个栈 duplicate 来维护，由于利润从大到小排序，
	所以栈顶就是最小的利润。入栈前判断 category 之前是否遇到过，遇到这入栈。
*/

func findMaximumElegance(items [][]int, k int) int64 {
	// 把利润从大到小排序
	slices.SortFunc(items, func(a, b []int) int { return b[0] - a[0] })
	ans, totalProfit := 0, 0
	vis := make(map[int]bool)
	var duplicate []int // 重复类别的利润
	for i, p := range items {
		profit, category := p[0], p[1]
		if i < k {
			totalProfit += profit // 累加前 k 个项目的利润
			if !vis[category] {
				vis[category] = true
			} else { // 重复类别
				duplicate = append(duplicate, profit)
			}
		} else if len(duplicate) > 0 && !vis[category] { // 之前没有的类别
			vis[category] = true                                // len(vis) 变大
			totalProfit += profit - duplicate[len(duplicate)-1] // // 选一个重复类别中的最小利润替换
			duplicate = duplicate[:len(duplicate)-1]
		} // else: 比前面的利润小，而且类别还重复了，选它只会让 totalProfit 变小，len(vis) 不变，优雅度不会变大
		ans = max(ans, totalProfit+len(vis)*len(vis))
	}
	return int64(ans)
}
