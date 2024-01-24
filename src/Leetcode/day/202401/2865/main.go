package main

func maximumSumOfHeights(a []int) int64 {
	ans := 0
	n := len(a)
	suf := make([]int, n+1)
	st := []int{n}
	sum := 0
	for i := n - 1; i >= 0; i-- {
		x := a[i]
		for len(st) > 1 && x <= a[st[len(st)-1]] {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			sum -= a[j] * (st[len(st)-1] - j)
		}
		sum += x * (st[len(st)-1] - i) // 从 i 到 st[len(st)-1]-1 都是x
		suf[i] = sum
		st = append(st, i)
	}
	ans = sum

	st = []int{-1}
	pre := 0
	for i, x := range a {
		for len(st) > 1 && x <= a[st[len(st)-1]] {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			pre -= a[j] * (j - st[len(st)-1])
		}
		pre += x * (i - st[len(st)-1])
		ans = max(ans, pre+suf[i+1])
		st = append(st, i)
	}

	return int64(ans)
}

func main() {

}
