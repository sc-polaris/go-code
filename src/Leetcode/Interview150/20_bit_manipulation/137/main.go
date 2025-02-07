package main

/*
	设只出现一次的那个数为 x。用二进制思考：
	如果 x 的某个比特是 0，由于其余数字都出现了 3 次，所以 nums 的所有元素在这个比特位上的 1 的个数是 3 的倍数。
	如果 x 的某个比特是 1，由于其余数字都出现了 3 次，所以 nums 的所有元素在这个比特位上的 1 的个数除 3 余 1。
*/

func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		cnt1 := int32(0)
		for _, x := range nums {
			cnt1 += int32(x) >> i & 1
		}
		ans |= cnt1 % 3 << i
	}
	return int(ans)
}

func singleNumber2(nums []int) int {
	a, b := 0, 0
	for _, x := range nums {
		b = (b ^ x) &^ a
		a = (a ^ x) &^ b
	}
	return b
}
