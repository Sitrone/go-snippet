package leetcode

// 给定条件为所有顶点的邻接边，dfs搜索 TODO 二次复习
func allPathsSourceTarget(graph [][]int) [][]int {
	ans := make([][]int, 0, len(graph))
	pathDfs(graph, []int{0}, &ans)
	return ans
}

func pathDfs(graph [][]int, path []int, ans *[][]int) {
	n := len(graph) - 1
	cur := path[len(path)-1]
	if cur == n {
		tmpPath := make([]int, len(path))
		copy(tmpPath, path)
		*ans = append(*ans, tmpPath)
		return
	}

	for _, node := range graph[cur] {
		path = append(path, node)
		pathDfs(graph, path, ans)
		path = path[:len(path)-1]
	}
}
