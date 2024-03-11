package main

import "fmt"

// & 运算符的优先级比 + -高
const (
	N   = 5010
	M   = 8192
	MOD = 1e9 + 7
)

var n int
var a [N]int
var f [2][M]int // 前i个数异或和为j的方案数

func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	f[0][0] = 1 // 0个数异或和为0的方案数为1
	for i := 1; i <= n; i++ {
		for j := 0; j < M; j++ {
			f[i&1][j] = f[(i-1)&1][j] // 不用第i个数
			if j^a[i] < M {
				f[i&1][j] = (f[i&1][j] + f[(i-1)&1][j^a[i]]) % MOD
			}
		}
	}

	var res int
	for i := 2; i < M; i++ {
		if isPrime(i) {
			res = (res + f[n&1][i]) % MOD
		}
	}

	fmt.Println(res)

}
