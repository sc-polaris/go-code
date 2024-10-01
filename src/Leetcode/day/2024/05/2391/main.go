package main

import "strings"

/*
	给你一个下标从 0 开始的字符串数组 garbage ，其中 garbage[i] 表示第 i 个房子的垃圾集合。garbage[i] 只包含字符 'M' ，'P' 和 'G' ，
	但可能包含多个相同字符，每个字符分别表示一单位的金属、纸和玻璃。垃圾车收拾 一 单位的任何一种垃圾都需要花费 1 分钟。

	同时给你一个下标从 0 开始的整数数组 travel ，其中 travel[i] 是垃圾车从房子 i 行驶到房子 i + 1 需要的分钟数。

	城市里总共有三辆垃圾车，分别收拾三种垃圾。每辆垃圾车都从房子 0 出发，按顺序 到达每一栋房子。但它们 不是必须 到达所有的房子。

	任何时刻只有 一辆 垃圾车处在使用状态。当一辆垃圾车在行驶或者收拾垃圾的时候，另外两辆车 不能 做任何事情。

	请你返回收拾完所有垃圾需要花费的 最少 总分钟数。
*/

/*
	总体思路：
	由于「任何时刻只有一辆垃圾车处在使用状态」，我们可以先让收集 M（金属）的垃圾车从左到右跑一趟，然后让收集 P（纸）的垃圾车从左到右跑跑一趟，
	最后让收集 G（玻璃）的垃圾车从左到右跑一趟。

	总用时可以拆分两个部分：
	1. 收集垃圾的总用时，这等于 garbage[i] 的长度之和。
	2. 行驶的总用时，这等于每辆垃圾车的行驶用时之和。对于收集 M 的垃圾车，设 garbage[i] 是最后一个包含 M 的字符串，那么收集 M 的垃圾车的行
	   驶用时为 travel[0] + travel[1] + ··· + travel[i-1]。同理可得另外两辆垃圾车的行驶用时。

	下面三种实现方法：
*/

/*
	方法一：多次遍历
	先把所有垃圾车都跑慢全程，再倒着遍历 garbage，减去多跑的时间。

	对于收集 M 的垃圾车，设 garbage[i] 是最后一个包含 M 的字符串，那么收集 M 的垃圾车的行驶用时，等于跑满全程的用时，减去多跑的用时，
	travel[i] + travel[i+1] + ··· + travel[n-2]，其中 n 是 garbage 的长度。注意 travel 的下标和 garbage 的下标相差 1。
*/

func garbageCollection(garbage []string, travel []int) (ans int) {
	for _, s := range garbage {
		ans += len(s)
	}
	for _, t := range travel {
		ans += t * 3
	}
	for _, c := range []byte("MPG") {
		for i := len(garbage) - 1; i > 0 && strings.IndexByte(garbage[i], c) < 0; i-- {
			ans -= travel[i-1] // 没有垃圾 c 多跑了
		}
	}
	return
}

/*
	方法二：正序一次遍历
	我们可以在遍历 garbage 的同时，计算从房子 0 到房子 i 的用时 sumT，以及一个哈希表（或者数组）tMap 记录每辆车目前的行驶用时。例如
	garbage[i] = GP，那么收集 G 和 P 的垃圾车的行驶用时更新为 sumT，收集 M 的垃圾车的行驶用时不变。循环结束后，tMap 中保存的就是每
	辆垃圾车各自的行驶用时了。

	最后答案为所有 garbage[i] 的长度之和，加上 tMap 中保存的行驶用时之和。
*/

func garbageCollection2(garbage []string, travel []int) int {
	ans := len(garbage[0])
	sumT := 0
	tMap := [4]int{}
	for i, g := range garbage[1:] {
		ans += len(g)
		sumT += travel[i]
		for _, c := range g {
			tMap[c&3] = sumT // MPG 的 ASCII 值的低两位互不相同
		}
	}
	for _, t := range tMap {
		ans += t
	}
	return ans
}

// 进一步地，在遍历 garbage 的过程中把行驶时间加入答案，从而做到一次遍历。

func garbageCollection3(garbage []string, travel []int) int {
	ans := len(garbage[0])
	sumT := 0
	tMap := [4]int{}
	for i, g := range garbage[1:] {
		ans += len(g)
		sumT += travel[i]
		for _, c := range g {
			ans += sumT - tMap[c&3]
			tMap[c&3] = sumT // MPG 的 ASCII 值的低两位互不相同
		}
	}
	return ans
}

/*
	方法三：倒序一次遍历（贡献法）
	方法二需要用到哈希表（hash map），能否只用哈希集合（hash set）呢？
	· 如果 garbage[n-1] 只包含 M，这意味着只有一辆车需要行驶 travel[n-2] 分钟，所以 travel[n-2] 对答案的贡献是 travel[n-2]·1
	· 如果 garbage[n-2]+garbage[n-1] 包含 M 和 P，这意味着有两辆车需要行驶 travel[n-3] 分钟，所以 travel[n-3] 对答案的贡献是 travel[n-3]·2
	· 以此类推，如果从 garbage[i] 到 garbage[n-1] 有 k 种字母，那么 travel[i-1] 对答案的贡献是 travel[i-1]·k。

	累加贡献即为行驶总用时。
*/

func garbageCollection4(garbage []string, travel []int) int {
	ans := len(garbage[0])
	seen := map[rune]struct{}{}
	for i := len(garbage) - 1; i > 0; i-- {
		g := garbage[i]
		for _, c := range g {
			seen[c] = struct{}{}
		}
		ans += len(g) + travel[i-1]*len(seen)
	}
	return ans
}
