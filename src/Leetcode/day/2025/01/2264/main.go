package main

import (
	"math"
	"strconv"
	"strings"
)

/*
	给你一个字符串 num ，表示一个大整数。如果一个整数满足下述所有条件，则认为该整数是一个 优质整数 ：
	· 该整数是 num 的一个长度为 3 的 子字符串 。
	· 该整数由唯一一个数字重复 3 次组成。
	以字符串形式返回 最大的优质整数 。如果不存在满足要求的整数，则返回一个空字符串 "" 。

	注意：
	· 子字符串 是字符串中的一个连续字符序列。
	· num 或优质整数中可能存在 前导零 。
*/

func largestGoodInteger(num string) (ans string) {
	ansNum := math.MinInt
	for i := 2; i < len(num); i++ {
		if num[i] == num[i-1] && num[i-1] == num[i-2] {
			tmp := num[i-2 : i+1]
			tmpNum, _ := strconv.Atoi(tmp)
			if ansNum < tmpNum {
				ansNum = tmpNum
				ans = tmp
			}
		}
	}
	return
}

func largestGoodInteger2(num string) string {
	mx := byte(0)
	cnt := 1
	for i := 1; i < len(num); i++ {
		d := num[i]
		if d != num[i-1] {
			cnt = 1
			continue
		}
		cnt++
		if cnt == 3 && d > mx {
			mx = d
		}
	}
	if mx == 0 {
		return ""
	}
	return strings.Repeat(string(mx), 3)
}
