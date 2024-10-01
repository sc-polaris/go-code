package main

/*
	如果 nums 的最后两个数相等，那么去掉这两个数，问题变成剩下 n-2 个数能否有效划分。
	如果 nums 的最后三个数相等，那么去掉这三个数，问题变成剩下 n-3 个数能否有效划分。
	如果 nums 的最后三个数是连续递增的，那么去掉这三个数，问题变成剩下 n-3 个数能否有效划分
	f[i] nums 的前 i 个数能否有效划分
	f[0] = true, f[i+1] 表示能否有效划分 nums[0] 到 nums[i]
*/

func validPartition(nums []int) bool {
	n := len(nums)
	f := make([]bool, n+1)
	f[0] = true
	for i, x := range nums {
		if i > 0 && f[i-1] && x == nums[i-1] {
			f[i+1] = true
		} else if i > 1 && f[i-2] {
			if x == nums[i-1] && x == nums[i-2] || x == nums[i-1]+1 && x == nums[i-2]+2 {
				f[i+1] = true
			}
		}
	}
	return f[n]
}

func main() {}
