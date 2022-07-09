package syncAdditional

import "sync"

/*
获取sync.map中读容器中的元素总数:
smap为指针
返回int类型
*/
func LenSyncMap(smap *sync.Map) int {
	count := 0
	smap.Range(func(key, value any) bool {
		count++
		return true
	})
	return count
}
