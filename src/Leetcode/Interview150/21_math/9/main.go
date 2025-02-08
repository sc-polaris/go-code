package main

/*
	反转一般数字
*/

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	r := 0
	for x > r {
		r = r*10 + x%10
		x /= 10
	}
	// 当数字长度为奇数时，我们可以通过 r/10 去除处于中位的数字。
	return x == r || x == r/10
}
