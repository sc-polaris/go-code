package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in   = bufio.NewReader(os.Stdin)
	ot   = bufio.NewWriter(os.Stdout)
	nums = make([]int, 0)
	ops  = make([]byte, 0)
	pr   = map[byte]int{'+': 1, '-': 1, '*': 2, '/': 2}
)

func eval() {
	b := nums[len(nums)-1]
	nums = nums[:len(nums)-1]
	a := nums[len(nums)-1]
	nums = nums[:len(nums)-1]
	c := ops[len(ops)-1]
	ops = ops[:len(ops)-1]

	var x int
	if c == '+' {
		x = a + b
	} else if c == '-' {
		x = a - b
	} else if c == '*' {
		x = a * b
	} else if c == '/' {
		x = a / b
	}

	nums = append(nums, x)
}

func isDigit(c byte) bool {
	if '0' <= c && c <= '9' {
		return true
	}
	return false
}

func main() {
	defer ot.Flush()

	var s string
	fmt.Fscan(in, &s)

	n := len(s)
	for i := 0; i < n; i++ {
		c := s[i]
		if isDigit(c) {
			v, j := 0, i
			for j < n && isDigit(s[j]) {
				v = v*10 + int(s[j]-'0')
				j++
			}
			i = j - 1
			nums = append(nums, v)
		} else if c == '(' {
			ops = append(ops, c)
		} else if c == ')' {
			for ops[len(ops)-1] != '(' {
				eval()
			}
			ops = ops[:len(ops)-1]
		} else {
			for len(ops) > 0 && pr[ops[len(ops)-1]] >= pr[c] {
				eval()
			}
			ops = append(ops, c)
		}
	}

	for len(ops) > 0 {
		eval()
	}

	fmt.Fprintln(ot, nums[len(nums)-1])
}
