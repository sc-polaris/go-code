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
