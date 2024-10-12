package main

/*
	给你一个数组 nums ，数组中的数字 要么 出现一次，要么 出现两次。

	请你返回数组中所有出现两次数字的按位 XOR 值，如果没有数字出现过两次，返回 0 。
*/

/*
	遍历 nums，同时用一个 vis 集合记录遇到的数字。
	1. 设 x=nums[i]。
	2. 如果 x 不在 vis 中，说明是第一次遇到，加入 vis。
	3. 如果 x 在 vis 中，说明是第二次遇到（注意每个数至多出现两次），加入答案（异或）。

	代码实现时，由于 nums[i]≤50，可以用二进制数表示集合，具体见
*/

func duplicateNumbersXOR(nums []int) (ans int) {
	vis := 0
	for _, x := range nums {
		if vis>>x&1 > 0 { // x 在 vis 中
			ans ^= x
		} else {
			vis |= 1 << x // 把 x 加到 vis 中
		}
	}
	return
}
