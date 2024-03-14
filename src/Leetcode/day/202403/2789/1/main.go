package main

func maxArrayValue(nums []int) int64 {
	n := len(nums)
	sum := nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] <= sum {
			sum += nums[i] // 继续向左合并
		} else {
			sum = nums[i]
		}
	}
	return int64(sum)
}

func main() {

}
