package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	n := len1 + len2
	i, j, l, r := 0, 0, 0, 0
	for count := 0; count <= n/2; count++ {
		l = r
		if i < len1 && (j >= len2 || nums1[i] < nums2[j]) {
			r = nums1[i]
			i++
		} else {
			r = nums2[j]
			j++
		}
	}
	if n&1 == 0 {
		return float64(l+r) / 2.0
	}
	return float64(r)
}
