package main

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	ans := -1
	t, mx := 0, 0 // t: 临时变量 mx最大利润
	wait, i := 0, 0
	for wait > 0 || i < len(customers) {
		if i < len(customers) {
			wait += customers[i]
		}
		up := min(4, wait)
		wait -= up
		t += up*boardingCost - runningCost
		i++
		if t > mx {
			mx = t
			ans = i
		}
	}
	return ans
}

func main() {

}
