package main

import "math"

/*
	给你两个 正整数 l 和 r。对于任何数字 x，x 的所有正因数（除了 x 本身）被称为 x 的 真因数。
	如果一个数字恰好仅有两个 真因数，则称该数字为 特殊数字。例如：
	· 数字 4 是 特殊数字，因为它的真因数为 1 和 2。
	· 数字 6 不是 特殊数字，因为它的真因数为 1、2 和 3。
	返回区间 [l, r] 内 不是 特殊数字 的数字数量。
*/

/*
	数据的最大范围为10^9，预处理sqrt(10^9)=31622
	然后用前缀和计算 [0,i] 中的质数个数
*/

const mx = 31622

var pi [mx + 1]int

func init() {
	for i := 2; i <= mx; i++ {
		if pi[i] == 0 {
			pi[i] = pi[i-1] + 1
			for j := i * i; j <= mx; j += i {
				pi[j] = -1 // 标记 i 的倍数为合数
			}
		} else {
			pi[i] = pi[i-1]
		}
	}
}

func nonSpecialCount(l int, r int) int {
	return r - l + 1 - (pi[int(math.Sqrt(float64(r)))] - pi[int(math.Sqrt(float64(l-1)))])
}
