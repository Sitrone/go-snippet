package leetcode

// 哈希表
func destCity(paths [][]string) string {
	startMap := make(map[string]struct{}, len(paths))
	for _, path := range paths {
		startMap[path[0]] = struct{}{}
	}

	for _, path := range paths {
		if _, ok := startMap[path[1]]; !ok {
			return path[1]
		}
	}

	return ""
}
