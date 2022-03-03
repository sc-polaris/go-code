package main

import "fmt"

const N int = 1010

var (
	n, m, T int
	a [N]int
)

func main() {
	fmt.Scanf("%d", &T)

	for ; T > 0; T-- {
        fmt.Scanf("%d %d", &n, &m)

		for i := 0; i < m; i++ {
			fmt.Scanf("%d", &a[i])
		}

		res := 0
		for i := 2; i <= n; i++ {
			res = (res + a[(n-2)%m]) % i
		}

		fmt.Println(res)
	}
}