package main

import "math"

func coinChange(coins []int, amount int) int {
	n := len(coins)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, amount+1)
	}
	for j := range f[0] {
		f[0][j] = math.MaxInt / 2 // 除 2 是防止下面 + 1 溢出
	}
	f[0][0] = 0
	for i, x := range coins {
		for c := 0; c <= amount; c++ {
			if c < x {
				f[i+1][c] = f[i][c]
			} else {
				f[i+1][c] = min(f[i][c], f[i+1][c-x]+1)
			}
		}
	}
	ans := f[n][amount]
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}

func main() {

}
