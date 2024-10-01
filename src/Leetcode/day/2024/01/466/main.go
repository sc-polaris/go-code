package main

// 循环优化
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	len1, len2 := len(s1), len(s2)
	index1, index2 := 0, 0
	if len1 == 0 || len2 == 0 || len1*n1 < len2*n2 {
		return 0
	}
	map1, map2 := make(map[int]int), make(map[int]int)
	ans := 0               // str1中s2的个数
	for index1/len1 < n1 { // 遍历整个str1
		if index1%len1 == len1-1 { // 在str1末尾
			if val, ok := map1[index2%len2]; ok { // 出现循环
				cycleLen := index1/len1 - val/len1                 // 每个循环占用多少个 s1
				cycleNum := (n1 - 1 - index1/len1) / cycleLen      // 还有多少个循环
				cycleS2Num := index2/len2 - map2[index2%len2]/len2 // 每个循环占多少个 s2

				index1 += cycleNum * cycleLen * len1 //将 index1 快进到相应的位置
				ans += cycleNum * cycleS2Num
			} else { // 第一次，注意存储的是未取模的
				map1[index2%len2] = index1
				map2[index2%len2] = index2
			}
		}

		if s1[index1%len1] == s2[index2%len2] {
			if index2%len2 == len2-1 {
				ans += 1
			}
			index2 += 1
		}
		index1 += 1
	}
	return ans / n2
}

func main() {

}
