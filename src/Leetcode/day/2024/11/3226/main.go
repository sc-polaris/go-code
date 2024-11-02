package main

import "math/bits"

/*
	给你两个正整数 n 和 k。

	你可以选择 n 的 二进制表示 中任意一个值为 1 的位，并将其改为 0。

	返回使得 n 等于 k 所需要的更改次数。如果无法实现，返回 -1
*/

// 如果 n 和 k 的交集是 k，那么 k 就是 n 的子集。 交集就是位运算中的 AND（&）。
func minChanges(n int, k int) int {
	if n&k != k {
		return -1
	}
	return bits.OnesCount(uint(n ^ k))
}

// 如果 n 和 k 的并集是 n，那么 k 就是 n 的子集。 并集就是位运算中的 OR（|）。
func minChanges2(n int, k int) int {
	if n|k != n {
		return -1
	}
	return bits.OnesCount(uint(n ^ k))
}

// 如果 k 去掉 n 中所有元素后，变成了空集，那么 k 就是 n 的子集。 写成代码，如果 (k & ~n) == 0，那么 k 就是 n 的子集。
func minChanges3(n, k int) int {
	if k&^n > 0 {
		return -1
	}
	return bits.OnesCount(uint(n ^ k))
}
