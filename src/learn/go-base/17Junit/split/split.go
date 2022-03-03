package split

import "strings"

// split package with a single split function.

// Split slices s into all substrings separated by sep and
// returns a slice of the substrings between those separators.

/*
func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)

	// 很显然我们最初的split函数并没有考虑到sep为多个字符的情况，
	for i > -1 {
		result = append(result, s[:i])
		//s = s[i+1:]
		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
*/

// Split 基准测试优化 这一次我们提前使用make函数将result初始化为一个容量足够大的切片，
// 而不再像之前一样通过调用append函数来追加。
func Split(s, sep string) (result []string) {
	result = make([]string, 0, strings.Count(s, sep)+1)
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
