package binary_index_tree

import (
	"testing"
)

func TestMaximumSumQueries(t *testing.T) {
	nums1 := []int{4, 3, 1, 2}
	nums2 := []int{2, 4, 9, 5}
	queries := [][]int{{4, 1}, {1, 3}, {2, 5}}
	t.Log(maximumSumQueries(nums1, nums2, queries))
}
