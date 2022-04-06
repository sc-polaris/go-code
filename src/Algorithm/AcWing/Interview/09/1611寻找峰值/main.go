package main

// Forward declaration of query API.

var nums [10010]int

func query(x int) int {
	return nums[x]
}

func findPeakElement(n int) int {
	l, r := 0, n-1
	for l < r {
		mid := (l + r) >> 1
		if query(mid) > query(mid+1) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return r
}

func main() {

}
