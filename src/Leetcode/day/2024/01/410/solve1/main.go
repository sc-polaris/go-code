package main

func splitArray(nums []int, k int) int {
	check := func(mx int) bool {
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
	}

	sum, mx := 0, 0
	for _, x := range nums {
		sum += x
		mx = max(mx, x)
	}
	left := max(mx, (sum-1)/k+1)
	right := sum
	for left < right {
		mid := left + (right-left)/2
		if !check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {

}
