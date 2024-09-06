package main

// 方法一：用哈希集合检测循环
func isHappy(n int) bool {
	step := func(n int) int {
		sum := 0
		for ; n > 0; n /= 10 {
			sum += (n % 10) * (n % 10)
		}
		return sum
	}
	m := make(map[int]bool)
	for ; n != 1 && !m[n]; n = step(n) {
		m[n] = true
	}
	return n == 1
}

// 方法二：快慢指针法
func isHappy2(n int) bool {
	step := func(n int) int {
		sum := 0
		for ; n > 0; n /= 10 {
			sum += (n % 10) * (n % 10)
		}
		return sum
	}
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}
