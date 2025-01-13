package main

/*
	给你三个 正 整数 num1 ，num2 和 num3 。

	数字 num1 ，num2 和 num3 的数字答案 key 是一个四位数，定义如下：
	· 一开始，如果有数字 少于 四位数，给它补 前导 0 。
	· 答案 key 的第 i 个数位（1 <= i <= 4）为 num1 ，num2 和 num3 第 i 个数位中的 最小 值。
	请你返回三个数字 没有 前导 0 的数字答案。
*/

func generateKey(x, y, z int) (ans int) {
	for pow10 := 1; x > 0 && y > 0 && z > 0; pow10 *= 10 {
		ans += min(x%10, y%10, z%10) * pow10
		x /= 10
		y /= 10
		z /= 10
	}
	return
}
