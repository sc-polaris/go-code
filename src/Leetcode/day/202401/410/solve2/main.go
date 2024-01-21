package main

import "sort"

func splitArray(nums []int, k int) int {
	sum, mx := 0, 0
	for _, x := range nums {
		sum += x
		mx = max(mx, x)
	}
	left := max(mx, (sum-1)/k+1)
	right := sum

	ans := left + sort.Search(right-left, func(mx int) bool {
		mx += left
		cnt, s := 1, 0
		for _, x := range nums {
			if s+x <= mx {
				s += x
			} else { // 新划分一段
				if cnt == k {
					return false
				}
				cnt += 1
				s = x
			}
		}
		return true
	})

	return ans
}

func main() {

}
