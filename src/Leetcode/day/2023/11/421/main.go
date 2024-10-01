package main

import (
	"math/bits"
	"slices"
)

func findMaximumXOR(nums []int) int {
	res := 0
	highBit := bits.Len(uint(slices.Max(nums))) - 1
	seen := make(map[int]bool)
	mask := 0
	for i := highBit; i >= 0; i-- {
		clear(seen)
		mask |= 1 << i       // 将第i位设位1
		newRes := res | 1<<i // 将 res第i位设为1
		for _, x := range nums {
			x &= mask // 低于i的比特位为 0
			if seen[newRes^x] {
				res = newRes
				break
			}
			seen[x] = true
		}
	}
	return res
}

func main() {

}
