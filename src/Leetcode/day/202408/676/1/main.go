package main

type MagicDictionary []string

func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (d *MagicDictionary) BuildDict(dictionary []string) {
	*d = dictionary
}

func (d *MagicDictionary) Search(searchWord string) bool {
next:
	for _, w := range *d {
		if len(w) != len(searchWord) {
			continue
		}
		diff := false
		for i := range w {
			if w[i] != searchWord[i] {
				if diff {
					continue next
				}
				diff = true
			}
		}
		if diff {
			return true
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
