package main

func maxResult(nums []int, k int) int {
	n := len(nums)
	f := make([]int, n)
	f[0] = nums[0]
	q := []int{0}
	for i := 1; i < n; i++ {
		// 1. 出
		if q[0] < i-k {
			q = q[1:]
		}
		// 2. 转移
		f[i] = f[q[0]] + nums[i]
		// 3. 入
		for len(q) > 0 && f[i] >= f[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	return f[n-1]
}

func main() {

}
