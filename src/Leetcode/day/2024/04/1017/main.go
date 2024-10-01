package main

/*
	n 从低位到高位的每一位，如果该位为 1，那么答案的该为为 1，否则为 0。如果该位为 1，我们需要将 n 减去 k。
	接下来我们更新 n /= 2，k = -k。继续判断下一位。

	最后答案反转

	3 = 011 = 111
	2 = 010 = 110
	4 = 100 = 100
*/

func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	var ans []byte
	k := 1
	for n != 0 {
		if n&1 == 1 {
			ans = append(ans, '1')
			n -= k
		} else {
			ans = append(ans, '0')
		}
		k *= -1
		n /= 2
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}
