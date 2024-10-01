package main

import "sort"

// 方法二：二分答案+数学公式

func findMaximumNumber(k int64, x int) int64 {
	ans := sort.Search(int(k+1)<<x, func(num int) bool {
		num++
		res := 0
		i := x - 1
		for n := num >> i; n > 0; n >>= x {
			res += n / 2 << i
			if n%2 > 0 {
				mask := 1<<i - 1
				res += num&mask + 1
			}
			i += x
		}
		return res > int(k)
	})
	return int64(ans)
}
