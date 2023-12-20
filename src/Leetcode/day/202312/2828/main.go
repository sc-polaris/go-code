package main

func isAcronym(words []string, s string) bool {
	var res []byte
	for _, word := range words {
		res = append(res, word[0])
	}
	return string(res) == s
}

func isAcronym2(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}
	for i := range s {
		if words[i][0] != s[i] {
			return false
		}
	}
	return true
}

func main() {

}
