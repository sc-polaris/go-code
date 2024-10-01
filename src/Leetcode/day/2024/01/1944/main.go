package main

func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	ans := make([]int, n)
	var st []int
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && st[len(st)-1] < heights[i] {
			st = st[:len(st)-1]
			ans[i]++
		}
		if len(st) > 1 { // 还可以再看到一个人
			ans[i]++
		}
		st = append(st, heights[i])
	}
	return ans
}

func main() {

}
