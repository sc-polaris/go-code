package main

/*
	从 i = k, j = k 出发，通过不断一动指针来找到最大矩形。比较 nums[i-1] 和 nums[j+1]
	的大小，谁大就移动谁（一样大的移动哪个都可以）
*/

func maximumScore(nums []int, k int) int {
	n := len(nums)
	ans, minH := nums[k], nums[k]
	i, j := k, k
	for t := 0; t < n-1; t++ { // 循环 n-1 次
		if j == n-1 || i > 0 && nums[i-1] > nums[j+1] {
			i--
			minH = min(minH, nums[i])
		} else {
			j++
			minH = min(minH, nums[j])
		}
		ans = max(ans, minH*(j-i+1))
	}
	return ans
}

func main() {

}
