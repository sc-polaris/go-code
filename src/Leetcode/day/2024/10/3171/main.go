package main

import "math"

/*
	给你一个数组 nums 和一个整数 k 。你需要找到 nums 的一个子数组，满足子数组中所有元素按位或运算 OR 的值与 k 的 绝对差 尽
	可能 小 。换言之，你需要选择一个子数组 nums[l..r] 满足 |k - (nums[l] OR nums[l + 1] ... OR nums[r])| 最小。

	请你返回 最小 的绝对差值。

	子数组 是数组中连续的 非空 元素序列。
*/

// 暴力 会超时
/*
func minimumDifference(nums []int, k int) int {
	ans := math.MaxInt
	for i, x := range nums {
		ans = min(ans, abs(x-k)) // 单个元素也算子数组
		for j := i - 1; j >= 0; j-- {
			nums[j] |= x // 现在 nums[j] = 原数组 nums[j] 到 nums[i] 的 OR
			ans = min(ans, abs(nums[j]-k))
		}
	}
	return ans
}
*/

// O(nlogU) 对于两个二进制数 a 和 b，如果 a ∣ b=a，那么 b 对应的集合是 a 对应的集合的子集。
func minimumDifference(nums []int, k int) int {
	ans := math.MaxInt
	for i, x := range nums {
		ans = min(ans, abs(x-k)) // 单个元素也算子数组
		for j := i - 1; j >= 0 && nums[j]|x != nums[j]; j-- {
			nums[j] |= x // 现在 nums[j] = 原数组 nums[j] 到 nums[i] 的 OR
			ans = min(ans, abs(nums[j]-k))
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
	方法二：滑动窗口 O(n)
	由于子数组越长，子数组 OR 的结果越大，有单调性，可以用滑动窗口。

	区别
	本题和普通滑窗的区别在于，不能只用一个变量 or 维护窗口（子数组）的 OR，因为当左端点元素离开窗口时，
	我们不知道要把 or 改成多少。本质上来说，是因为 OR 不像加法，没有逆运算（加法的逆运算是减法）。

	额外维护一个栈
	例如现在窗口的下标为 left=0 到 right=3。当左端点元素 nums[0] 离开窗口时，我们必须有一个值能够表示 nums[1] 到 nums[3] 的 OR。

	这里的思路是：
	· nums[3] 不变，它表示下标从 3 到 3 这个子数组的 OR。
	· 把 nums[2] 更新为 nums[2]∣nums[3]。
	· 把 nums[1] 更新为 nums[1]∣nums[2]，也就是原数组的 nums[1]∣nums[2]∣nums[3]。
	· 想象有一个栈，栈底是 nums[3]，栈顶是 nums[1]。
	现在窗口元素的 OR，即 nums[1] 到 nums[3] 的 OR，就存储在栈顶 nums[1] 中了。如果又有一个左端点元素离开窗口，我们就弹出栈顶，用新的栈顶 nums[2] 表示窗口元素的 OR。

	当右端点 right 移动时，我们又该如何维护呢？
	额外维护一个变量 rightOr，表示从 nums[4] 到 nums[right] 的 OR。

	现在窗口元素分成了两部分：
	· 左部：nums[left] 到 nums[3]。这些元素的 OR 保存在 nums[left] 中。
	· 右部：nums[4] 到 nums[right]。这些元素的 OR 保存在 rightOr 中。

	那么计算 nums[left] 和 rightOr 的 OR，就是窗口元素的 OR 了。
	但问题又来了，如果我们把栈清空了，也就是左端点移动到 left=4，又该怎么做呢？
	重复之前的过程，从 i=right−1 开始，倒序循环到 left，更新 nums[i] 为 nums[i]∣nums[i+1]。相当于我们在计算一个后缀 OR。这个过程结束后，相当于生成了一个栈，栈底是
	nums[right]，栈顶是 nums[left]。每个栈中元素 nums[i] 表示从 nums[i] 到 nums[right] 的 OR。计算完后，把 rightOr 重置为 0。

	为了记录栈底的位置，我们还需要维护一个额外变量 bottom。

	更新答案
	原问题可以拆分成（等价于）：
	· 当窗口元素 OR 大于 k 时，计算窗口元素 OR 的最小值。
	· 当窗口元素 OR 小于等于 k 时，计算窗口元素 OR 的最大值。

	所以本题既是一个求最小值的滑窗，又是一个求最大值的滑窗。
	具体来说：
	· 右端点元素进入窗口后，如果发现 (nums[left]∣rightOr)>k，那么用 (nums[left]∣rightOr)−k 更新答案的最小值。
	· 左端点移动结束后（退出内层循环），现在 (nums[left]∣rightOr)≤k，那么用 k−(nums[left]∣rightOr) 更新答案的最小值。
*/

func minimumDifference2(nums []int, k int) int {
	ans := math.MaxInt
	var left, bottom, rightOr int
	for right, x := range nums {
		rightOr |= x
		for left <= right && nums[left]|rightOr > k {
			ans = min(ans, (nums[left]|rightOr)-k)
			if bottom <= left {
				// 重新构建一个栈
				// 由于 left 即将移出窗口，只需计算到 left+1
				for i := right - 1; i > left; i-- {
					nums[i] |= nums[i+1]
				}
				bottom = right
				rightOr = 0
			}
			left++
		}
		if left <= right {
			ans = min(ans, k-(nums[left]|rightOr))
		}
	}
	return ans
}
