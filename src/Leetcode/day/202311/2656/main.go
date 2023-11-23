package main

import "slices"

/*
m + (m + 1) + (m + 2) + ... + (m + k -1)
等差数列求和
an=a1+(n-1)d
Sn=a1n+n(n-1)d/2
Sn=m*k+k(k-1)/2=(2m+k-1)/2
*/

func maximizeSum(nums []int, k int) int {
	return (2*slices.Max(nums) + k - 1) * k / 2
}
