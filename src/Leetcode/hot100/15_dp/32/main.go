package main

func longestValidParentheses(s string) int {
	maxAns := 0
	st := []int{-1}
	for i, c := range s {
		if c == '(' {
			st = append(st, i)
		} else {
			st = st[:len(st)-1]
			if len(st) == 0 {
				st = append(st, i)
			} else {
				maxAns = max(maxAns, i-st[len(st)-1])
			}
		}
	}
	return maxAns
}
