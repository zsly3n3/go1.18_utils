package slice

// ObverseDeleteSlice 从正序开始删除指定元素
func ObverseDeleteSlice[T any](a []T, index int) []T {
	j := 0
	for k, v := range a {
		if k != index {
			a[j] = v
			j++
		}
	}
	return a[:j]
}
