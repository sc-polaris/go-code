package main

/*

	摩尔投票法（Boyer–Moore majority vote algorithm），也被称作「多数投票法」，算法解决的问题是：如何在任意多的候选人中（选票无序），选出获得票数最多的那个。

	算法可以分为两个阶段：
	1. 对抗阶段：分属两个候选人的票数进行两两对抗抵消
	2. 计数阶段：计算对抗结果中最后留下的候选人票数是否有效
*/

func majorityElement(nums []int) int {
	major, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			major = v
		}
		if major == v {
			count++
		} else {
			count--
		}
	}
	return major
}
