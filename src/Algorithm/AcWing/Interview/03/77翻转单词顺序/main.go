package main

import "strings"

// 1

func reverseWords(s string) string {
	tmp := []byte(s)
	reverse(tmp, 0, len(tmp))

	for i := 0; i < len(tmp); i++ {
		j := i
		for j < len(s) && tmp[j] != ' ' {
			j++
		}
		reverse(tmp, i, j)
		i = j
	}

	return string(tmp)
}

func reverse(s []byte, l, r int) {
	r--
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

// 2
func solve2(s string) string {
	n := len(s)
	res := ""
	// 先对两头trim，确定左右边界
	l, r := 0, n-1
	for l < n && s[l] == ' ' {
		l++
	}
	for r >= 0 && s[r] == ' ' {
		r--
	}
	j := r + 1
	for i := r; i >= l; i-- {
		if s[i] == ' ' {
			res = res + s[i+1:j] + " "
			// 如果左边还是空格就继续
			for i >= 1 && s[i-1] == ' ' {
				i--
			}
			// j记录单词后的第一个空格位置
			j = i
		} else if i == l {
			res += s[i:j]
		}
	}
	return res
}

// 3
func solve3(s string) string {
	// Fields：按空格将字符串切分成字符串切片
	res := strings.Fields(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return strings.Join(res, " ")
}
