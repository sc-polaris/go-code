package _375__二进制字符串前缀一致的次数

func numTimesAllBlue(flips []int) int {
	n, res, right := len(flips), 0, 0
	for i := 0; i < n; i++ {
		right = max(flips[i], right)
		if right == i+1 {
			res++
		}
	}
	return res
}
