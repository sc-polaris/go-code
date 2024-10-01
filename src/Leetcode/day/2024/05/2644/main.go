package main

import "slices"

/*
	给你两个下标从 0 开始的整数数组 nums 和 divisors 。

	divisors[i] 的 可整除性得分 等于满足 nums[j] 能被 divisors[i] 整除的下标 j 的数量。

	返回 可整除性得分 最大的整数 divisors[i] 。如果有多个整数具有最大得分，则返回数值最小的一个。

*/

// 暴力 O(nm)
func maxDivScore(nums []int, divisors []int) (ans int) {
	maxCnt := -1
	for _, d := range divisors {
		cnt := 0
		for _, x := range nums {
			if x%d == 0 {
				cnt++
			}
		}
		if cnt > maxCnt || cnt == maxCnt && d < ans {
			maxCnt, ans = cnt, d
		}
	}
	return
}

/*
	优化 1
	注意到，小于 d 的正整数无法被 d 整除
	把 nums 排序，从大到小遍历 nums，只需遍历 >= d 的 nums[i]，当 nums[i] < d 时，退出内层循环

	O(nlogn + nm)
*/

func maxDivScore2(nums []int, divisors []int) (ans int) {
	slices.SortFunc(nums, func(a, b int) int { return b - a })
	maxCnt := -1
	for _, d := range divisors {
		cnt := 0
		for _, x := range nums {
			if x < d {
				break
			}
			if x%d == 0 {
				cnt++
			}
		}
		if cnt > maxCnt || cnt == maxCnt && d < ans {
			maxCnt, ans = cnt, d
		}
	}
	return
}

/*
	优化 2
	在优化 1 的基础上，把 divisors 从小到大排序，并统计 nums 中重复元素的个数 dup，例如 nums = [3,3,3,2,1,1]，其中右 dup=3 个数时重复的。
	遍历 d = divisors[i]，如果
							(maxCnt - dup + 1)*d > max(nums)
	说明 d 的倍数 d,2d,3d,···,(maxCnt - dup + 1)*d 中的最大值已经超出了 nums 的最大值，即使把 nums 中的重复元素也算上，我们也无法统计出比
	maxCnt 还多的倍数。由于我们已经把 divisors 从小到大排序了，当前的 d 不满足上面的不等式，那后面更大的 d 也同样满足上面的不等式，所以后面不可
	能找到一个比 maxCnt 更大的数，直接退出外部循环。

	避免乘法溢出，改为判断
							maxCnt - dup + 1 > ⌊max(nums)/d⌋
	即
							maxCnt - dip >= ⌊max(nums)/d⌋

	O(nlogn + mlogm + nm)
*/

func maxDivScore3(nums []int, divisors []int) (ans int) {
	slices.SortFunc(nums, func(a, b int) int { return b - a })
	dup := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			dup++
		}
	}
	slices.Sort(divisors)
	maxCnt := -1
	for _, d := range divisors {
		if maxCnt-dup > nums[0]/d {
			break
		}
		cnt := 0
		for _, x := range nums {
			if x < d {
				break
			}
			if x%d == 0 {
				cnt++
			}
		}
		if cnt > maxCnt {
			maxCnt, ans = cnt, d
		}
	}
	return
}
