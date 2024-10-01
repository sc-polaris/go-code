package main

func minCostClimbingStairs(cost []int) int {
	f0, f1 := 0, 0
	for i := 1; i < len(cost); i++ {
		f0, f1 = f1, min(f1+cost[i], f0+cost[i-1])
	}
	return f1
}

func main() {

}
