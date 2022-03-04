package main

import "fmt"

const (
	N   = 1e5 + 10
	INF = 0x3f3f3f3f
)

var n int
var a [N]int

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	maxProfit, minPrice := 0, INF
	for i := 0; i < n; i++ {
		maxProfit = max(maxProfit, a[i]-minPrice)
		minPrice = min(minPrice, a[i])
	}

	fmt.Println(maxProfit)
}
