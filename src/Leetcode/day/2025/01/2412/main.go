package main

/*

	给你一个下标从 0 开始的二维整数数组 transactions，其中transactions[i] = [costi, cashbacki] 。

	数组描述了若干笔交易。其中每笔交易必须以 某种顺序 恰好完成一次。在任意一个时刻，你有一定数目的钱 money ，
	为了完成交易 i ，money >= costi 这个条件必须为真。执行交易后，你的钱数 money 变成 money - costi + cashbacki 。

	请你返回 任意一种 交易顺序下，你都能完成所有交易的最少钱数 money 是多少。
*/

/*
	「任意一种交易顺序下，都能完成所有交易」意味着要考虑在最坏情况下，需要多少初始钱数 initMoney。
	什么是最坏情况？
	先亏钱（cost>cashback），再赚钱（cost≤cashback），主打一个欲扬先抑。
	初始钱数必须满足，在最穷困潦倒的时候，也能完成交易。
	什么时候最穷？完成所有亏钱交易后最穷。

	记 totalLose 为所有亏钱的 cost−cashback 之和。
	遍历 transactions，分类讨论：
	· 对于赚钱的交易，假设这是（亏钱后的）第一笔赚钱的交易，那么初始钱数是多少？为了完成这笔交易，题目
	  要求此时的钱至少是 cost，即 initMoney−totalLose≥cost，得 initMoney≥totalLose+cost。
	· 对于亏钱的交易，假设这是最后一笔亏钱的交易，那么初始钱数是多少？由于 cost−cashback 已经计入
	  totalLose 中，需要先从 totalLose 中减去 cost−cashback，即
				initMoney−(totalLose−(cost−cashback))≥cost，
	  化简得到
				initMoney≥totalLose+cashback。
	所有情况取最大值，就能保证在任意一种交易顺序下，都能完成所有交易。
	· 如果赚钱，即 cost≤cashback，那么 totalLose 加上的是二者的较小值 cost。
	· 如果亏钱，即 cost>cashback，那么 totalLose 加上的也是二者的较小值 cashback。
	综上所述，初始钱数 initMoney 等于 totalLose 加上 min(cost,cashback) 的最大值。
*/

func minimumMoney(transactions [][]int) int64 {
	totalLose, mx := 0, 0
	for _, t := range transactions {
		totalLose += max(t[0]-t[1], 0)
		mx = max(mx, min(t[0], t[1]))
	}
	return int64(totalLose + mx)
}
