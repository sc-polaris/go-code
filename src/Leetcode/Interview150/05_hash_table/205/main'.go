package main

func isIsomorphic(s string, t string) bool {
	s2t := make(map[byte]byte)
	t2s := make(map[byte]byte)
	for i := range s {
		x, y := s[i], t[i]
		if yy, ok := s2t[x]; ok && yy != y {
			return false
		}
		if xx, ok := t2s[y]; ok && xx != x {
			return false
		}
		s2t[x] = y
		t2s[y] = x
	}
	return true
}
