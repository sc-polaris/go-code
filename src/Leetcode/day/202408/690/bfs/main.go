package bfs

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) (res int) {
	eMap := make(map[int]*Employee, len(employees))
	for _, e := range employees {
		eMap[e.Id] = e
	}

	q := []int{id}
	for len(q) > 0 {
		tmp := eMap[q[0]]
		q = q[1:]
		res += tmp.Importance
		for _, subId := range tmp.Subordinates {
			q = append(q, subId)
		}
	}
	return
}
