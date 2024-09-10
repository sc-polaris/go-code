package main

import (
	"math"
	"strings"
)

func calculate(s string) (ans int) {
	m := map[byte]int{'+': 1, '-': 1, '*': 2, '/': 2, '%': 2, '^': 3}
	nums := []int{0}
	var ops []byte
	s = strings.ReplaceAll(s, " ", "")
	n := len(s)

	// 计算函数
	calc := func() {
		if len(nums) < 2 || len(ops) == 0 {
			return
		}
		num1, num2 := nums[len(nums)-2], nums[len(nums)-1]
		nums = nums[:len(nums)-2]
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]
		ans := 0
		switch op {
		case '+':
			ans = num1 + num2
		case '-':
			ans = num1 - num2
		case '*':
			ans = num1 * num2
		case '/':
			ans = num1 / num2
		case '%':
			ans = num1 % num2
		case '^':
			ans = int(math.Pow(float64(num1), float64(num2)))
		}
		nums = append(nums, ans)
	}

	for i := 0; i < n; i++ {
		c := s[i]
		switch c {
		case '(':
			ops = append(ops, c)
		case ')':
			for len(ops) > 0 {
				if ops[len(ops)-1] != '(' {
					calc()
				} else {
					ops = ops[:len(ops)-1]
					break
				}
			}
		default:
			if '0' <= c && c <= '9' {
				num := 0
				j := i
				for ; j < n && '0' <= s[j] && s[j] <= '9'; j++ {
					num = num*10 + int(s[j]-'0')
				}
				i = j - 1
				nums = append(nums, num)
			} else {
				if i > 0 && (s[i-1] == '(' || s[i-1] == '+' || s[i-1] == '-') {
					nums = append(nums, 0)
				}
				for len(ops) > 0 && ops[len(ops)-1] != '(' && m[ops[len(ops)-1]] >= m[c] {
					calc()
				}
				ops = append(ops, c)
			}
		}
	}

	for len(ops) > 0 && ops[len(ops)-1] != '(' {
		calc()
	}

	return nums[len(nums)-1]
}
