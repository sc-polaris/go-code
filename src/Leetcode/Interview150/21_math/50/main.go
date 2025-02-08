package main

func myPow(x float64, n int) float64 {
	ans := 1.0
	if n < 0 {
		n = -n
		x = 1 / x
	}
	for n > 0 {
		if n&1 == 1 {
			ans *= x
		}
		x *= x
		n >>= 1
	}
	return ans
}
