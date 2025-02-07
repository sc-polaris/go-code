package main

func hammingWeight(n int) (res int) {
	for n > 0 {
		n &= n - 1
		res++
	}
	return
}
