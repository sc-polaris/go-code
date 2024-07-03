package main

/*
	如果一个整数能够被其各个数位上的数字之和整除，则称之为 哈沙德数（Harshad number）。给你一个整数 x 。如果 x 是 哈沙德数 ，
	则返回 x 各个数位上的数字之和，否则，返回 -1 。
*/

func sumOfTheDigitsOfHarshadNumber(x int) int {
	s := 0
	for t := x; t > 0; t /= 10 {
		s += t % 10
	}
	if x%s == 0 {
		return s
	}
	return -1
}
