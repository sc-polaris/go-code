package dfs

/*

	你有一个保存员工信息的数据结构，它包含了员工唯一的 id ，重要度和直系下属的 id 。

	给定一个员工数组 employees，其中：
	· employees[i].id 是第 i 个员工的 ID。
	· employees[i].importance 是第 i 个员工的重要度。
	· employees[i].subordinates 是第 i 名员工的直接下属的 ID 列表。
	给定一个整数 id 表示一个员工的 ID，返回这个员工和他所有下属的重要度的 总和。

*/

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	eMap := make(map[int]*Employee, len(employees))
	for _, e := range employees {
		eMap[e.Id] = e
	}
	var dfs func(int) int
	dfs = func(id int) int {
		e := eMap[id]
		res := e.Importance
		for _, subId := range e.Subordinates {
			res += dfs(subId)
		}
		return res
	}
	return dfs(id)
}
