package main

func canConstruct(ransomNote string, magazine string) bool {
	cnt := [26]int{}
	for _, c := range magazine {
		cnt[c-'a']++
	}
	for _, c := range ransomNote {
		if cnt[c-'a'] == 0 {
			return false
		}
		cnt[c-'a']--
	}
	return true
}
