package main

/*
	有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。

	现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。
	这里的 i - 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。

	求所能获得硬币的最大数量。
*/

/*
	我们记数组 nums 的长度为 n。根据题目描述，我们可以在数组 nums 的左右两端各添加一个 1，记为 arr。

	f[i][j] 表示戳破区间 [i,j] 内的所有气球能得到的最多硬币数，那么答案即为 f[0][n+1]

	对于 f[i][j]，我们枚举区间 [i][j] 内的所有位置 k，假设 k 是最后一个戳破的气球，那么我们可以得到如下状态转移方程：
			f[i][j] = max(f[i][j], f[i][k]+f[k][j]+arr[i]*arr[k]*arr[j])

*/

func maxCoins(nums []int) int {
	n := len(nums)
	arr := make([]int, n+2)
	arr[0] = 1
	arr[n+1] = 1
	copy(arr[1:n+1], nums)

	f := make([][]int, n+2)
	for i := range f {
		f[i] = make([]int, n+2)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				f[i][j] = max(f[i][j], f[i][k]+f[k][j]+arr[i]*arr[k]*arr[j])
			}
		}
	}

	return f[0][n+1]
}
