package main

import (
	"slices"
)

/*
	给定数组 people 。people[i]表示第 i 个人的体重 ，船的数量不限，每艘船可以承载的最大重量为 limit。

	每艘船最多可同时载两人，但条件是这些人的重量之和最多为 limit。

	返回 承载所有人所需的最小船数 。
*/

func numRescueBoats(people []int, limit int) (ans int) {
	slices.Sort(people)
	l, r := 0, len(people)-1
	for l <= r {
		if people[l]+people[r] > limit {
			r--
		} else {
			l++
			r--
		}
		ans++
	}
	return
}
