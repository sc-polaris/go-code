package main

func canCompleteCircuit(gas []int, cost []int) int {
	allSum, run, start := 0, 0, 0
	for i := range gas {
		run += gas[i] - cost[i]
		allSum += gas[i] - cost[i]
		if run < 0 {
			start = i + 1
			run = 0
		}
	}
	if allSum < 0 {
		return -1
	}
	return start
}
