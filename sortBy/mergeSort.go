package sortBy

import (
	"fmt"
	"github.com/zsly3n3/go1.18_utils/common"
)

/*
数值归并排序,可升可降的排:
asc为true是升序,false则为降序
arr为目标排序数组也是排序结果
start,end为排序范围,索引范围
*/
func MergeSortNumeric[T common.NumericValue](asc bool, arr []T, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	MergeSortNumeric(asc, arr, start, mid)
	MergeSortNumeric(asc, arr, mid+1, end)
	mergeNumeric(asc, arr, start, mid, end)
}

func mergeNumeric[T common.NumericValue](asc bool, arr []T, start, mid, end int) {
	tmpArr := make([]T, 0, end-start+1)
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
func MergeSortFloat64Struct(asc bool, arr []MergeSortFloat64Obj, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	MergeSortFloat64Struct(asc, arr, start, mid)
	MergeSortFloat64Struct(asc, arr, mid+1, end)
	mergeFloat64Struct(asc, arr, start, mid, end)
}

func mergeFloat64Struct(asc bool, arr []MergeSortFloat64Obj, start, mid, end int) {
	tmpArr := make([]MergeSortFloat64Obj, 0, end-start+1)
	var s1, s2 = start, mid + 1
	for s1 <= mid && s2 <= end {
		if asc {
			if arr[s1].GetFloatValue() > arr[s2].GetFloatValue() {
				tmpArr = append(tmpArr, arr[s2])
				s2++
			} else {
				tmpArr = append(tmpArr, arr[s1])
				s1++
			}
		} else {
			if arr[s1].GetFloatValue() < arr[s2].GetFloatValue() {
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

type MergeSortFloat64Obj interface {
	GetFloatValue() float64
}

func MergeSortIntStruct(asc bool, arr []MergeSortIntObj, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	MergeSortIntStruct(asc, arr, start, mid)
	MergeSortIntStruct(asc, arr, mid+1, end)
	mergeIntStruct(asc, arr, start, mid, end)
}

func mergeIntStruct(asc bool, arr []MergeSortIntObj, start, mid, end int) {
	tmpArr := make([]MergeSortIntObj, 0, end-start+1)
	var s1, s2 = start, mid + 1
	for s1 <= mid && s2 <= end {
		if asc {
			if arr[s1].GetIntValue() > arr[s2].GetIntValue() {
				tmpArr = append(tmpArr, arr[s2])
				s2++
			} else {
				tmpArr = append(tmpArr, arr[s1])
				s1++
			}
		} else {
			if arr[s1].GetIntValue() < arr[s2].GetIntValue() {
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

type MergeSortIntObj interface {
	GetIntValue() int
}

/**
 * 递归算法实现二分查找
 *
 * @param nums  数组
 * @param low   左下标
 * @param high  右下标
 * @param value 要查找的值
 * @return
 */
func recursiveBserach(nums []int, low, high, value int) int {
	if low > high {
		return -1
	}
	// 找出中间下标
	mid := low + ((high - low) >> 1)
	fmt.Println("mid:", mid)
	if nums[mid] == value {
		return mid
	} else if nums[mid] > value {
		return recursiveBserach(nums, low, mid-1, value)
	} else {
		return recursiveBserach(nums, mid+1, high, value)
	}
}
