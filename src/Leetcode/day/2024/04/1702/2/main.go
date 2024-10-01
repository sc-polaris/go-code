package main

/*
	操作1： 00->10	操作2：10->01
	模拟：
	我们从字符串左边第一位开始依次遍历，如果是 1 则不用改变，如果是 0，我们则想办法将其变成 1。

	我们会找到下一位出现的 0，利用操作 2 我们可以使得这两个 0 相邻，再使用操作 1 使得 00 变成 10。

	我们依次执行这个操作，直到字符串中没有第二个 0，或者达到字符串结尾。
*/

func maximumBinaryString(binary string) string {
	n := len(binary)
	s := []rune(binary)
	j := 0
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			for j <= i || (j < n && s[j] == '1') {
				j++
			}
			if j < n {
				s[j] = '1'
				s[i] = '1'
				s[i+1] = '0'
			}
		}
	}
	return string(s)
}
