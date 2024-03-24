package main

import "math"

/*
	空间优化：两个数组（滚动数组）
*/

func coinChange(coins []int, amount int) int {
	n := len(coins)
	f := make([][]int, 2)
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
				f[(i+1)%2][c] = f[i%2][c]
			} else {
				f[(i+1)%2][c] = min(f[i%2][c], f[(i+1)%2][c-x]+1)
			}
		}
	}
	ans := f[n%2][amount]
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}

func main() {

}
