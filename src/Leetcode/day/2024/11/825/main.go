package main

/*
	在社交媒体网站上有 n 个用户。给你一个整数数组 ages ，其中 ages[i] 是第 i 个用户的年龄。

	如果下述任意一个条件为真，那么用户 x 将不会向用户 y（x != y）发送好友请求：
	· ages[y] <= 0.5 * ages[x] + 7
	· ages[y] > ages[x]
	· ages[y] > 100 && ages[x] < 100
	否则，x 将会向 y 发送一条好友请求。

	注意，如果 x 向 y 发送一条好友请求，y 不必也向 x 发送一条好友请求。另外，用户不会向自己发送好友请求。

	返回在该社交媒体网站上产生的好友请求总数。
*/

func numFriendRequests(ages []int) (ans int) {
	cnt := [121]int{}
	for _, age := range ages {
		cnt[age]++
	}

	cntWindow, y := 0, 0
	for x, c := range cnt[:] {
		cntWindow += c
		if y*2 <= x+14 { // 不能发送好友请求
			cntWindow -= cnt[y]
			y++
		}
		if cntWindow > 0 { // 存在可以发送好友请求的用户
			ans += c*cntWindow - c
		}
	}
	return
}
