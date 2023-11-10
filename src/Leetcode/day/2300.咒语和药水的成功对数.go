package day

import (
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	m := len(potions)
	res := make([]int, len(spells))
	for i, v := range spells {
		res[i] = m - sort.Search(m, func(i int) bool { return int64(potions[i]*v) >= success })
	}
	return res
}
