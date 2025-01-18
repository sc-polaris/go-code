package main

import "math"

/*
	给你一个整数数组 nums 和一个 正 整数 k 。

	定义长度为 2 * x 的序列 seq 的 值 为：
	· (seq[0] OR seq[1] OR ... OR seq[x - 1]) XOR (seq[x] OR seq[x + 1] OR ... OR seq[2 * x - 1]).
	请你求出 nums 中所有长度为 2 * k 的子序列的 最大值 。
*/

/*
	题意
	从 nums 中选一个长为 2k 的子序列，计算其前一半的 OR，后一半的 OR，这两个 OR 再计算 XOR。
	问：计算出的 XOR 最大能是多少？

	核心思路
	· 想象有一根分割线，把 nums 分成左右两部分，左和右分别计算所有长为 k 的子序列的 OR 都有哪些值。比如左边计算出的 OR
	  有 2,3,5，右边计算出的 OR 有 1,3,6，那么两两组合计算 XOR，其中最大值即为答案。
	· 枚举分割线的位置，把 nums 分割成一个前缀和一个后缀，问题变成：从前缀/后缀中选一个长为 k 的子序列，这个子序列 OR
	  的结果能否等于 x？

	把 OR 理解成一个类似加法的东西，转换成二维 0-1 背包。
	二维：指背包有两个约束，一个是所选元素的个数是 k，另一个是所选元素的 OR 是 x。

	具体算法
	计算后缀。对于 0-1 背包问题，我们定义 f[i][j][x] 表示从 nums[i] 到 nums[n−1] 中选 j 个数，这些数的 OR 能否等于 x。
	设 v=nums[i]，用刷表法转移：
	· 不选 v，那么 f[i][j][x]=f[i+1][j][x]。
	· 选 v，如果 f[i+1][j][x]=true，那么 f[i][j+1][x|v]=true。

	刷表法：本题计算 x=v|? 中的 ? 是困难的，但计算 x|v 是很简单的。也就是说，对于状态 f[i][j][x] 而言，其转移来源是谁不好计算，
	但从 f[i][j][x] 转移到的目标状态 f[i][j+1][x|v] 是好计算的。在动态规划中，根据转移来源计算状态叫查表法，用当前状态更新其他
	状态叫刷表法。

	初始值 f[n][0][0]=true。什么也不选，OR 等于 0。
	对于每个 i，由于我们只需要 f[i][k] 中的数据，把 f[i][k] 复制到 suf[i] 中。这样做无需创建三维数组，空间复杂度更小。
	代码实现时，f 的第一个维度可以优化掉。
	对于前缀 pre 的计算也同理。
	最后，枚举 i=k−1,k,k+1,…,n−k−1，两两组合 pre[i] 和 suf[i+1] 中的数计算 XOR，其中最大值即为答案。
	小优化：如果在循环中，发现答案 ans 达到了理论最大值 2^7−1（或者所有元素的 OR），则立刻返回答案。

	也可以用哈希集合代替布尔数组，见下面的 Python 优化代码。
*/

func maxValue(nums []int, k int) (ans int) {
	const mx = 1 << 7
	n := len(nums)
	suf := make([][mx]bool, n-k+1)
	f := make([][mx]bool, k+1)
	f[0][0] = true
	for i := n - 1; i >= k; i-- {
		v := nums[i]
		// 注意当 i 比较大的时候，循环次数应和 i 有关，因为更大的 j，对应的 f[j] 全为 false
		for j := min(k-1, n-1-i); j >= 0; j-- {
			for x, hasX := range f[j] {
				if hasX {
					f[j+1][x|v] = true
				}
			}
		}
		if i <= n-k {
			suf[i] = f[k]
		}
	}

	clear(f)
	f[0][0] = true
	for i, v := range nums[:n-k] {
		for j := min(k-1, i); j >= 0; j-- {
			for x, hasX := range f[j] {
				if hasX {
					f[j+1][x|v] = true
				}
			}
		}
		if i < k-1 {
			continue
		}
		// 这里 f[k] 就是 pre[i]
		for x, hasX := range f[k] {
			if hasX {
				for y, hasY := range suf[i+1] {
					if hasY {
						ans = max(ans, x^y)
					}
				}
			}
		}
		if ans == mx-1 {
			return
		}
	}
	return
}

// 优化后

const bitWidth = 7
const mx = 1 << bitWidth

func maxValue2(nums []int, k int) (ans int) {
	n := len(nums)
	k2 := min(k, bitWidth) // 至多选 k2 个数
	suf := make([][mx]bool, n-k+1)
	f := make([][mx]bool, k2+1)
	f[0][0] = true
	maxI := [mx]int{}
	cnt := [mx]int{}
	for i := n - 1; i >= k; i-- {
		v := nums[i]
		for j := min(k2-1, n-1-i); j >= 0; j-- {
			for x, hasX := range f[j] {
				if hasX {
					f[j+1][x|v] = true
				}
			}
		}
		if i <= n-k {
			suf[i] = f[k2]
		}
		// 枚举 v 的超集
		for s := v; s < mx; s = (s + 1) | v {
			cnt[s]++
			if cnt[s] == k {
				// 从 n-1 开始遍历，至少要遍历到 i 才有可能找到 k 个数 OR 等于 s
				maxI[s] = i
			}
		}
	}

	pre := make([][mx]bool, k2+1)
	pre[0][0] = true
	minI := [mx]int{}
	for i := range minI {
		minI[i] = math.MaxInt
	}
	cnt = [mx]int{}
	for i, v := range nums[:n-k] {
		for j := min(k2-1, i); j >= 0; j-- {
			for x, hasX := range pre[j] {
				if hasX {
					pre[j+1][x|v] = true
				}
			}
		}
		// 枚举 v 的超集
		for s := v; s < mx; s = (s + 1) | v {
			cnt[s]++
			if cnt[s] == k {
				// 从 0 开始遍历，至少要遍历到 i 才有可能找到 k 个数 OR 等于 s
				minI[s] = i
			}
		}
		if i < k-1 {
			continue
		}
		var a []int
		var b []int
		for x, has := range pre[k2] {
			if has && minI[x] <= i {
				a = append(a, x)
			}
			if suf[i+1][x] && maxI[x] > i {
				b = append(b, x)
			}
		}
		ans = max(ans, findMaximumXOR(a, b))
		if ans == mx-1 {
			return
		}
	}
	return
}

// 421. 数组中两个数的最大异或值
// 改成两个数组的最大异或值，做法是类似的，仍然可以用【试填法】解决
func findMaximumXOR(a, b []int) (ans int) {
	mask := 0
	for i := bitWidth - 1; i >= 0; i-- { // 从最高位开始枚举
		mask |= 1 << i
		newAns := ans | 1<<i // 这个比特位可以是 1 吗？
		seen := [mx]bool{}
		for _, x := range a {
			seen[x&mask] = true // 低于 i 的比特位置为 0
		}
		for _, x := range b {
			x &= mask // 低于 i 的比特位置为 0
			if seen[newAns^x] {
				ans = newAns
				break
			}
		}
	}
	return
}
