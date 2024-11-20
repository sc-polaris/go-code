package main

func minMutation(start string, end string, bank []string) int {
	if start == end {
		return 0
	}
	bankSet := map[string]bool{}
	for _, s := range bank {
		bankSet[s] = true
	}
	type pair struct {
		gene string
		step int
	}
	q := []pair{{start, 0}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		cur := p.gene
		for i, x := range cur {
			for _, y := range "ACGT" {
				if y != x {
					nxt := cur[:i] + string(y) + cur[i+1:]
					if bankSet[nxt] {
						if nxt == end {
							return p.step + 1
						}
						bankSet[nxt] = false
						q = append(q, pair{nxt, p.step + 1})
					}
				}
			}
		}
	}
	return -1
}
