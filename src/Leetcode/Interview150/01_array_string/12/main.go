package main

import "fmt"

// 硬编码

var R = [4][10]string{
	{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}, // 个位
	{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}, // 十位
	{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}, // 百位
	{"", "M", "MM", "MMM"}, // 千位
}

func intToRoman(num int) (ans string) {
	return R[3][num/1000] + R[2][num/100%10] + R[1][num/10%10] + R[0][num%10]
}

// 模拟
var valueSymbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman2(num int) (ans string) {
	for _, vs := range valueSymbols {
		for num >= vs.value {
			num -= vs.value
			ans = fmt.Sprintf("%s%s", ans, vs.symbol)
		}
		if num == 0 {
			break
		}
	}
	return
}
