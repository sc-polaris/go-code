package main

import (
	"slices"
	"sort"
)

/*
	计算 nums 中所有非负数的和，记作 sum
	nums 的任意一个子序列的元素和，都等价于从 sum 中减去某些非负数/加上某些负数得到。
	解决问题转换：
		把每个 nums[i] 取绝对值后，nums 的第 k 小的子序列和是多少？
	二分答案，设当前二分的值为 sumLimit。
	问题变成：判断是否有至少 k 个子序列，元素和 s 不超过 sumLimit。
	注：一道题能否二分答案，得看它有没有单调性。对于本题，sumLimit 越大，这样的子序列越多，有单调性，可以二分答案。

	问：有没有可能，二分得到的值，并不是 nums 的子序列和？比如 nums[i] 都是偶数，但二分得到的却是一个奇数
	答：设二分得到的值是为 x，那么 x 一定是 nums 的子序列和。使用反证法证明：
		假设 x 不是 nums 的子序列和，也就是没有任何子序列和等于 x，这一意味着 s<=x 等价于 s<=x-1，我们能从
		nums 中找到 k 个元素和不超过 x-1 的子序列，所以 check(x-1)=true。但二分循环结束时，有 check(x-1)=false
		矛盾，所以原命题成立，x 一定是 nums 的子序列和。
*/

func kSum(nums []int, k int) int64 {
	sum, total := 0, 0
	for i, x := range nums {
		if x >= 0 {
			sum += x
			total += x
		} else {
			total -= x
			nums[i] = -x
		}
	}
	slices.Sort(nums)

	kthS := sort.Search(total, func(sumLimit int) bool {
		cnt := 1 // 空子序列算第一个
		var dfs func(int, int)
		dfs = func(i int, s int) {
			if cnt == k || i == len(nums) || s+nums[i] > sumLimit {
				return
			}
			cnt++               // s + nums[i] <= sumLimit
			dfs(i+1, s+nums[i]) // 选
			dfs(i+1, s)         // 不选
		}
		dfs(0, 0)
		return cnt == k // 找到 k 个元素和不超过 sumLimit 的子序列
	})
	return int64(sum - kthS)
}

func main() {}
