package main

import (
	"math"
	"sort"
)

/*

a1 = n(n+1)/2
k=n+1 d=n
Sk = a1k + k(k-1)d/2
   = [n(n+1)^2+(n+1)n^2]/2
   = n(n+1)(2n+1)/2

正方形中苹果数量 = 4Sk = 2n(n+1)(2n+1)
正方形周长 = 8n
求最小的 n
		2n(n+1)(2n+1) >= neededApples 得
		 n(n+1)(2n+1) >= neededApples/4
设	m = cbrt(neededApples/4)
因为		(m-1)m(m-1/2)<m^3<=neededApples/4
所以 m-1 必不满足
因为		(m+1)(m+2)(m+3/2)>(m+1)^3>neededApples/4
所以 m+1 必满足要求
即 m+1 > cbrt(neededApples/4)
*/

func minimumPerimeter(neededApples int64) int64 {
	n := int64(math.Cbrt(float64(neededApples) / 4))
	if 2*n*(n+1)*(2*n+1) < neededApples {
		n++
	}
	return 8 * n
}

func minimumPerimeter2(neededApples int64) int64 {
	ans := sort.Search(100000, func(i int) bool {
		return int64(2*i*(i+1)*(2*i+1)) >= neededApples
	})

	return 8 * int64(ans)
}

func main() {

}
