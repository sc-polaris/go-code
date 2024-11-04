package main

import "math"

/*
	给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a^2 + b^2 = c 。
*/

/*
	数学：
	费马平方和定理告诉我们：一个非负整数 c 如果能够表示为两个整数的平方和，当且仅当 c 的所有形如 4k+3 的质因子的幂均为偶数。
*/

func judgeSquareSum(c int) bool {
	for base := 2; base*base <= c; base++ {
		if c%base != 0 {
			continue
		}

		// 计算 base 的幂
		exp := 0
		for c%base == 0 {
			c /= base
			exp++
		}

		if base%4 == 3 && exp%2 != 0 {
			return false
		}
	}
	// 退出循环以后需要再做一次判断
	return c%4 != 3
}

// 枚举
func judgeSquareSum2(c int) bool {
	for a := 0; a*a <= c/2; a++ {
		b := int(math.Sqrt(float64(c - a*a)))
		if a*a+b*b == c {
			return true
		}
	}
	return false
}

// 相向双指针
func judgeSquareSum3(c int) bool {
	a, b := 0, int(math.Sqrt(float64(c)))
	for a <= b {
		s := a*a + b*b
		if s == c {
			return true
		}
		if s < c {
			a++
		} else {
			b--
		}
	}
	return false
}
