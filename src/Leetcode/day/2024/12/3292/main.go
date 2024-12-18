package main

/*
	给你一个字符串数组 words 和一个字符串 target。

	如果字符串 x 是 words 中 任意 字符串的前缀，则认为 x 是一个 有效 字符串。

	现计划通过 连接 有效字符串形成 target ，请你计算并返回需要连接的 最少 字符串数量。如果无法通过这种方式形成 target，则返回 -1。
*/

/*
	1. Z函数
	示例 1 的 words=[abc,aaaaa,bcdef]，target=aabcdabc。
	要让划分的段数尽量小，那么每一段的长度要尽量长？
	虽然还不知道要怎么划分，但这启发我们考虑如下内容：
	· 从 target[0] 开始的段，最长可以是多长？答案是 2，因为 aa 是 words[1]=aaaaa 的前缀。
	· 从 target[1] 开始的段，最长可以是多长？答案是 3，因为 abc 是 words[0]=abc 的前缀。
	· 从 target[2] 开始的段，最长可以是多长？答案是 3，因为 bcd 是 words[2]=bcdef 的前缀
	· ……
	注意 words[i] 前缀的前缀还是 words[i] 的前缀。bcd 是 bcdef 的前缀，意味着 b 和 bc 也是 bcdef 的前缀。如果从 target[2] 开始的段，最长可以是 3，那么这个段的长度也可以是 1 或者 2。
	如果我们算出了上面这些最长长度 2,3,3,⋯，那么问题就变成：
	· 给你一个数组 maxJumps=[2,3,3,⋯]。你从 0 开始向右跳。在下标 i 处，可以跳到 [i+1,i+maxJumps[i]] 中的任意下标。目标是到达 n，最小要跳多少次？
	示例 1 的 maxJumps=[2,3,3,0,0,3,2,0]，跳法是 0→2→5→8。
	现在剩下的问题是，如何计算 maxJumps 数组？

	对于字符串 s，定义 z[i] 表示后缀 s[i:] 与 s 的 LCP（最长公共前缀）长度，其中 s[i:] 表示从 s[i] 到 s[n−1] 的子串。
	遍历 words，对于 word=words[i]，构造字符串
							s=word+#+target
	中间插入 # 目的是避免 z[i] 超过 word 的长度。
	计算 s 的 z 数组。设 m 为 word 的长度加一，那么 target[i:] 与 word 的最长公共前缀，就是 z[m+i]。用 z[m+i] 更新 maxJumps[i] 的最大值。
*/

func minValidStrings(words []string, target string) int {
	calcZ := func(s string) []int {
		n := len(s)
		z := make([]int, n)
		boxL, boxR := 0, 0 // z-box 左右边界（闭区间）
		for i := 1; i < n; i++ {
			if i <= boxR {
				z[i] = min(z[i-boxL], boxR-i+1)
			}
			for i+z[i] < n && s[z[i]] == s[i+z[i]] {
				boxL, boxR = i, i+z[i]
				z[i]++
			}
		}
		return z
	}
	jump := func(maxJumps []int) (ans int) {
		curR := 0 // 已建造的桥的右端点
		nxtR := 0 // 下一座桥的右端点的最大值
		for i, maxJump := range maxJumps {
			nxtR = max(nxtR, i+maxJump)
			if i == curR { // 到达已建造的桥的右端点
				if i == nxtR { // 无论怎么造桥，都无法从 i 到 i+1
					return -1
				}
				curR = nxtR // 造一座桥
				ans++
			}
		}
		return
	}

	maxJumps := make([]int, len(target))
	for _, word := range words {
		z := calcZ(word + "#" + target)
		for i, z := range z[len(word)+1:] {
			maxJumps[i] = max(maxJumps[i], z)
		}
	}
	return jump(maxJumps)
}
