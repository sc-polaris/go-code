package main

func uniqueLetterString(s string) (res int) {
	// last0: 上一次出现的下标 last1: 上上一次出现的下标
	last0, last1, total := [26]int{}, [26]int{}, 0
	for i := range last0 {
		last0[i] = -1
		last0[i] = -1
	}

	for i, c := range s {
		c -= 'A'
		total += (i - last0[c]) - (last0[c] - last1[c])
		res += total
		last1[c] = last0[c]
		last0[c] = i
	}

	return res
}

func main() {

}
