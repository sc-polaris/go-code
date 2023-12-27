package main

import "fmt"

func isWinner(player1 []int, player2 []int) int {
	cal := func(player []int) int {
		res := 0
		for i, v := range player {
			if i > 0 && player[i-1] == 10 || i > 1 && player[i-2] == 10 {
				res += 2 * v
			} else {
				res += v
			}
		}
		return res
	}

	res1, res2 := cal(player1), cal(player2)

	if res1 > res2 {
		return 1
	} else if res1 < res2 {
		return 2
	} else {
		return 0
	}
}

func main() {
	fmt.Println(isWinner([]int{5, 6, 1, 10}, []int{5, 1, 10, 5}))
}
