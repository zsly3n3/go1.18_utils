package sortBy

type numericValue interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

/*
数值归并排序,可升可降的排:
asc为true是升序,则为降序
arr为目标排序数组也是排序结果
start,end为排序范围
*/
func MergeSortNumeric[T numericValue](asc bool, arr []T, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	MergeSortNumeric(asc, arr, start, mid)
	MergeSortNumeric(asc, arr, mid+1, end)
	mergeNumeric(asc, arr, start, mid, end)
}

func mergeNumeric[T numericValue](asc bool, arr []T, start, mid, end int) {
	var tmpArr []T
	var s1, s2 = start, mid + 1
	for s1 <= mid && s2 <= end {
		if asc {
			if arr[s1] > arr[s2] {
				tmpArr = append(tmpArr, arr[s2])
				s2++
			} else {
				tmpArr = append(tmpArr, arr[s1])
				s1++
			}
		} else {
			if arr[s1] < arr[s2] {
				tmpArr = append(tmpArr, arr[s2])
				s2++
			} else {
				tmpArr = append(tmpArr, arr[s1])
				s1++
			}
		}
	}
	if s1 <= mid {
		tmpArr = append(tmpArr, arr[s1:mid+1]...)
	}
	if s2 <= end {
		tmpArr = append(tmpArr, arr[s2:end+1]...)
	}
	for pos, item := range tmpArr {
		arr[start+pos] = item
	}
}

/*
结构体归并排序,可升可降的排:
asc为true是升序,则为降序
arr为目标排序数组也是排序结果,需要实现MergeSortObj接口
start,end为排序范围
*/
func MergeSortStruct(asc bool, arr []MergeSortObj, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	MergeSortStruct(asc, arr, start, mid)
	MergeSortStruct(asc, arr, mid+1, end)
	mergeStruct(asc, arr, start, mid, end)
}

func mergeStruct(asc bool, arr []MergeSortObj, start, mid, end int) {
	var tmpArr []MergeSortObj
	var s1, s2 = start, mid + 1
	for s1 <= mid && s2 <= end {
		if asc {
			if arr[s1].GetValue() > arr[s2].GetValue() {
				tmpArr = append(tmpArr, arr[s2])
				s2++
			} else {
				tmpArr = append(tmpArr, arr[s1])
				s1++
			}
		} else {
			if arr[s1].GetValue() < arr[s2].GetValue() {
				tmpArr = append(tmpArr, arr[s2])
				s2++
			} else {
				tmpArr = append(tmpArr, arr[s1])
				s1++
			}
		}
	}
	if s1 <= mid {
		tmpArr = append(tmpArr, arr[s1:mid+1]...)
	}
	if s2 <= end {
		tmpArr = append(tmpArr, arr[s2:end+1]...)
	}
	for pos, item := range tmpArr {
		arr[start+pos] = item
	}
}

type MergeSortObj interface {
	GetValue() float64
}
