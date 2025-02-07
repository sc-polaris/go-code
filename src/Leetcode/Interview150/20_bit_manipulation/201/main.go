package main

import "math/bits"

func rangeBitwiseAnd(left int, right int) int {
	m := bits.Len(uint(left ^ right))
	return left &^ ((1 << m) - 1)
}
