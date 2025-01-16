package main

import "math"

/*
	给你一个 非负 整数数组 nums 和一个整数 k 。

	如果一个数组中所有元素的按位或运算 OR 的值 至少 为 k ，那么我们称这个数组是 特别的 。

	请你返回 nums 中 最短特别非空子数组的长度，如果特别子数组不存在，那么返回 -1 。
*/

/*
	LogTrick
	首先，我们有如下 O(n^2) 的暴力算法：

	外层循环，从 i=0 开始，从左到右遍历 nums。内层循环，从 j=i−1 开始，从右到左遍历 nums，更新 nums[j]=nums[j]|nums[i]。
	· i=1 时，我们会把 nums[0] 到 nums[1] 的 OR 记录在 nums[0] 中。
	· i=2 时，我们会把 nums[1] 到 nums[2] 的 OR 记录在 nums[1] 中，nums[0] 到 nums[2] 的 OR 记录在 nums[0] 中。
	· i=3 时，我们会把 nums[2] 到 nums[3] 的 OR 记录在 nums[2] 中；nums[1] 到 nums[3] 的 OR 记录在 nums[1] 中；nums[0] 到 nums[3] 的 OR 记录在 nums[0] 中。
	· 依此类推。
	按照该算法，可以计算出所有子数组的 OR。注意单个元素也算子数组。

	对于两个二进制数 a 和 b，如果 a|b=a，那么 b 对应的集合是 a 对应的集合的子集
*/

func minimumSubarrayLength(nums []int, k int) int {
	ans := math.MaxInt
	for i, x := range nums {
		if x >= k {
			return 1
		}
		for j := i - 1; j >= 0 && nums[j]|x != nums[j]; j-- {
			nums[j] |= x
			if nums[j] >= k {
				ans = min(ans, i-j+1)
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

/*
	滑动窗口+栈
	由于子数组越长，子数组 OR 的结果越大，有单调性，可以用滑动窗口。

	本题和普通滑窗的区别在于，不能只用一个变量 or 维护窗口（子数组）的 OR，因为当左端点元素离开窗口时，我们不知道要把 or 改成多少。本质上来说，是因为 OR 不像加法，没有逆运算（加法的逆运算是减法）。

	额外维护一个栈
	例如现在窗口的下标为 left=0 到 right=3。当左端点元素 nums[0] 离开窗口时，我们必须有一个值能够表示 nums[1] 到 nums[3] 的 OR。
	这里的思路是：
	· nums[3] 不变，它表示下标从 3 到 3 这个子数组的 OR。
	· 把 nums[2] 更新为 nums[2]|nums[3]。
	· 把 nums[1] 更新为 nums[1]|nums[2]，也就是原数组的 nums[1]|nums[2]|nums[3]。
	· 想象有一个栈，栈底是 nums[3]，栈顶是 nums[1]。

	现在窗口元素的 OR，即 nums[1] 到 nums[3] 的 OR，就存储在栈顶 nums[1] 中了。如果又有一个左端点元素离开窗口，我们就弹出栈顶，用新的栈顶 nums[2] 表示窗口元素的 OR。
	当右端点 right 移动时，我们又该如何维护呢？
	额外维护一个变量 rightOr，表示从 nums[4] 到 nums[right] 的 OR。
	现在窗口元素分成了两部分：
	· 左部：nums[left] 到 nums[3]。这些元素的 OR 保存在 nums[left] 中。
	· 右部：nums[4] 到 nums[right]。这些元素的 OR 保存在 rightOr 中。

	那么计算 nums[left] 和 rightOr 的 OR，就是窗口元素的 OR 了。
	但问题又来了，如果我们把栈清空了，也就是左端点移动到 left=4，又该怎么做呢？

	重复之前的过程，从 i=right−1 开始，倒序循环到 left，更新 nums[i] 为 nums[i]|nums[i+1]。相当于我们在计算一个后缀 OR。这个过程结束后，相当于生成了一个栈，栈底是
	nums[right]，栈顶是 nums[left]。每个栈中元素 nums[i] 表示从 nums[i] 到 nums[right] 的 OR。计算完后，把 rightOr 重置为 0。

	为了记录栈底的位置，我们还需要维护一个额外变量 bottom。

	更新答案
	原问题可以拆分成（等价于）：
	· 当窗口元素 OR 大于 k 时，计算窗口元素 OR 的最小值。
	· 当窗口元素 OR 小于等于 k 时，计算窗口元素 OR 的最大值。
*/

func minimumSubarrayLength2(nums []int, k int) int {
	ans := math.MaxInt
	var left, bottom, rightOr int
	for right, x := range nums {
		rightOr |= x
		for left <= right && nums[left]|rightOr >= k {
			ans = min(ans, right-left+1)
			left++
			if bottom < left {
				// 重新构建一个栈
				for i := right - 1; i >= left; i-- {
					nums[i] |= nums[i+1]
				}
				bottom = right
				rightOr = 0
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
