package main

func maxProfit(prices []int) (ans int) {
	pre := prices[0]
	for _, v := range prices[1:] {
		ans = max(ans, v-pre)
		pre = min(pre, v)
	}
	return
}
