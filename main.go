package main

import (
	"fmt"
	"github.com/zsly3n3/go1.18_utils/generate"
	"golang.org/x/exp/constraints"
)

func main() {
	//arr := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	//fmt.Println(slice.ReverseDeleteSlice(arr, 5))
}

func Test[T constraints.Ordered](arr []T) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

func Abc() {
	var ts uint64 = 1861545531 //高33位
	var region uint64 = 3      //低4位
	n := 100000
	mp := make(map[uint64]struct{})
	for i := 1; i <= n; i++ {
		id := uint64(i)
		guid := generate.CreateUserGuid(ts, id, region)
		_, ok := mp[guid]
		if ok {
			fmt.Println(`isExist`)
			break
		} else {
			mp[guid] = struct{}{}
		}
	}
	for k := range mp {
		ts0, _, region0 := generate.GetDataFromGuid(k)
		if ts0 != ts {
			fmt.Println(`ts error`)
			break
		}
		if region0 != region {
			fmt.Println(`region error`)
			break
		}
	}
	fmt.Println(`---over`)
}
