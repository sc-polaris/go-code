package main

func isValid(s string) bool {
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
