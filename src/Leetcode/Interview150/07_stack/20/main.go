package main

// 写法一
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	mp := map[rune]rune{')': '(', ']': '[', '}': '{'}
	var st []rune
	for _, c := range s {
		if mp[c] == 0 { // c 是左括号
			st = append(st, c)
		} else { // c 是右括号
			if len(st) == 0 || st[len(st)-1] != mp[c] {
				return false // 没有左括号，或者左括号类型不对
			}
			st = st[:len(st)-1] // 出栈
		}
	}
	return len(st) == 0
}

// 写法二
func isValid2(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	mp := map[rune]rune{'(': ')', '[': ']', '{': '}'}
	var st []rune
	for _, c := range s {
		if v := mp[c]; v > 0 { // c 是左括号
			st = append(st, v)
		} else { // c 是右括号
			if len(st) == 0 || st[len(st)-1] != c {
				return false // 没有左括号，或者左括号类型不对
			}
			st = st[:len(st)-1] // 出栈
		}
	}
	return len(st) == 0
}

// 写法三
func isValid3(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	var st []rune
	for _, c := range s {
		switch c {
		case '(':
			st = append(st, ')')
		case '[':
			st = append(st, ']')
		case '{':
			st = append(st, '}')
		default: // c 是右括号
			if len(st) == 0 || st[len(st)-1] != c {
				return false
			}
			st = st[:len(st)-1]
		}
	}
	return len(st) == 0
}
