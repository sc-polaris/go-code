package main

func nextBeautifulNumber(n int) int {
	isBalance := func(x int) bool {
		cnt := make([]int, 10)
		for x > 0 {
			cnt[x%10]++
			x /= 10
		}
		for i := 0; i < 10; i++ {
			if cnt[i] > 0 && cnt[i] != i {
				return false
			}
		}
		return true
	}

	for x := n + 1; ; x++ {
		if isBalance(x) {
			return x
		}
	}
}

func main() {

}
