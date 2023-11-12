package day

import (
	"testing"
)

func Test(t *testing.T) {
	spells := []int{5, 1, 3}
	potions := []int{1, 2, 3, 4, 5}
	res := successfulPairs(spells, potions, 7)
	t.Log(res)
}

func Test2(t *testing.T) {
	// 0 1 0 1
	row := []int{0, 2, 1, 3}
	t.Log(minSwapsCouples(row))
}

func TestRangeModule(t *testing.T) {
	obj := Constructor()
	obj.AddRange(10, 20)
	obj.RemoveRange(14, 16)
	t.Log(obj.QueryRange(10, 14))
	t.Log(obj.QueryRange(13, 15))
	t.Log(obj.QueryRange(16, 17))
}
