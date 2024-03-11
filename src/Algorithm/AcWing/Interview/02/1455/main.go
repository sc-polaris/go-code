package main

import "fmt"

// 1455. 招聘

/*
约瑟夫环递推公式：
    f(1) = 0;     //表示最后一轮的胜出者当前编号是0
    f(x) = (f(x - 1) + m) % x , 1 < x <= n //每一轮都找到胜出者在上一轮中的编号
    不过本题里m是在变化的，所以要相应地变为：
==> f(x) = (f(x - 1) + a[(n - x) % m]) % x, 1 < x <= n
*/

/*
当只有一个人时，答案res = 0，我们从两个人开始递推，这里需要注意我们的真实操作从
n个人变为n-1个人用的是a[0]，从n-1个人变为n-2个人我们用的是a[1]，则从2个人变为1
个人我们用的是a[(n-2)%m]
*/

const N int = 1010

var (
	T    int
	n, m int
	a    [N]int
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
			res = (res + a[(n-i)%m]) % i
		}

		fmt.Println(res)
	}
}
