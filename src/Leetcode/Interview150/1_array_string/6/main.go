package main

/*
	1. 以 V 字型为一个循环, 循环周期为 n = (2 * numRows - 2) （2倍行数 - 头尾2个）。
	2. 对于字符串索引值 i，计算 x = i % n 确定在循环周期中的位置。
	3. 则行号 y = min(x, n - x)。
*/

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]string, numRows)
	n := 2*numRows - 2
	for i, c := range s {
		x := i % n
		rows[min(x, n-x)] += string(c)
	}
	return strings.Join(rows, "")
}
