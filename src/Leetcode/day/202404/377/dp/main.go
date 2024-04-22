package dp

func combinationSum4(nums []int, target int) int {
	f := make([]int, target+1)
	f[0] = 1
	for i := 1; i <= target; i++ {
		for _, x := range nums {
			if x <= i {
				f[i] += f[i-x]
			}
		}
	}
	return f[target]
}
