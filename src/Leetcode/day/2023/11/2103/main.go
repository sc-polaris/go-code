package main

import "fmt"

func countPoints(rings string) int {
	state := make([]int, 10)
	n := len(rings)
	for i := 0; i < n; i += 2 {
		color := rings[i]
		index := rings[i+1] - '0'
		if color == 'R' {
			state[index] |= 1
		} else if color == 'G' {
			state[index] |= 2
		} else {
			state[index] |= 4
		}
	}

	res := 0
	for i := 0; i < 10; i++ {
		if state[i] == 7 {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(countPoints("B0B6G0R6R0R6G9"))
}
