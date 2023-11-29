package main

const l = 10

var bin = map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}

func findRepeatedDnaSequences(s string) (res []string) {
	n := len(s)
	if n <= l {
		return
	}

	// 二进制20个比特计算10的子串
	x := 0
	for _, ch := range s[:l-1] {
		x = x<<2 | bin[byte(ch)]
	}

	cnt := make(map[int]int)
	for i := 0; i <= n-l; i++ {
		// 最左边的字符离开窗口 x = x & ((1<<20)-1)
		x = (x<<2 | bin[s[i+l-1]]) & (1<<(l*2) - 1)
		cnt[x]++
		if cnt[x] == 2 {
			res = append(res, s[i:i+l])
		}
	}
	return
}

func main() {

}
