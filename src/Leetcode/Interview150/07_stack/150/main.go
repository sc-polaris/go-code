package main

import "strconv"

func evalRPN(tokens []string) int {
	var st []int
	for _, token := range tokens {
		if val, err := strconv.Atoi(token); err == nil {
			st = append(st, val)
		} else {
			num1, num2 := st[len(st)-2], st[len(st)-1]
			st = st[:len(st)-2]
			switch token {
			case "+":
				st = append(st, num1+num2)
			case "-":
				st = append(st, num1-num2)
			case "*":
				st = append(st, num1*num2)
			case "/":
				st = append(st, num1/num2)
			}
		}
	}
	return st[0]
}
