package main

import "math"

func buyChoco(prices []int, money int) int {
	a, b := math.MaxInt, math.MaxInt
	for _, x := range prices {
		if x < a {
			a, b = x, a
		} else if x < b {
			b = x
		}
	}

	if money < a+b {
		return money
	}
	return money - a - b
}

func main() {

}
