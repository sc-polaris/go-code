package main

import "strconv"

func maximumSwap(num int) int {
	s := strconv.Itoa(num)
	maxIdx := len(s) - 1
	p, q := -1, 0
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] > s[maxIdx] { // s[i] 是目前最大数字
			maxIdx = i
		} else if s[i] < s[maxIdx] { // s[i] 右边有比它大的
			p, q = i, maxIdx // 更新 p 和 q
		}
	}
	if p == -1 { // 说明 s 是降序的
		return num
	}
	t := []byte(s)
	t[p], t[q] = t[q], t[p]
	ans, _ := strconv.Atoi(string(t))
	return ans
}

func main() {

}
