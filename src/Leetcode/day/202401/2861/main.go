package main

import (
	"slices"
	"sort"
)

/*
	假如制造 num 份金合金
	if composition[i][j]*num <= stock[j] 无需购买额外的金属
	if composition[i][j]*num >  stock[j] 需要购买额外的金属，花费为
		(composition[i][j]*num-stock[j])*cost[j]
	二分上界：
		假设 composition[i][j] 和 cost[j] 都是 1，此时可以制造最多的合金，个数为
			min(stock) + budget
	二分下界：
		可以设为 1。更巧妙的做法是，设当前答案为 ans，那么可以初始化二分下界为 ans+1，
		因为酸楚小于等于 ans 的值是没有意义的，不会让 ans 变大。如果这台机器无法制造出
		至少 ans+1 份合金，那么二分循环结束后的结果为 ans，不影响答案
	左闭有开区间写法 go模板库也是这个写法
	上界+1
*/

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) (ans int) {
	mx := slices.Max(stock) + budget
	for _, comp := range composition {
		ans += sort.Search(mx-ans, func(num int) bool {
			num += ans + 1
			money := 0
			for i, s := range stock {
				if s < comp[i]*num {
					money += (comp[i]*num - s) * cost[i]
					if money > budget {
						return true
					}
				}
			}
			return false
		})
	}
	return
}

func maxNumberOfAlloys2(n int, k int, budget int, composition [][]int, stock []int, cost []int) (ans int) {
	mx := slices.Max(stock) + budget
	for _, comp := range composition {
		check := func(num int) bool {
			money := 0
			for i, s := range stock {
				if s < comp[i]*num {
					money += (comp[i]*num - s) * cost[i]
					if money > budget {
						return true
					}
				}
			}
			return false
		}
		l, r := 0, mx-ans
		for l < r {
			mid := int(uint(l+r)) >> 1
			if !check(mid + ans + 1) {
				l = mid + 1
			} else {
				r = mid
			}
		}
		ans += l
	}
	return
}

func main() {

}
