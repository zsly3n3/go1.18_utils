package main

import (
	"fmt"
	"github.com/zsly3n3/go1.18_utils/generate"
	"golang.org/x/exp/constraints"
)

func main() {
	//s1 := make([]float64, 0)
	//s1 = append(s1, 5, 3, 2)
	//s1 = append(s1, 1.0, 0, 9)
	//s1 = append(s1, 2.5, 8, 11)
	//num := len(s1)
	//sortBy.MergeSortNumeric(false, s1, 0, num-1)
	//Test(s1)
	//
	//smap := sync.Map{}
	//smap.Store(`t1`, 1)
	//count := syncAdditional.LenSyncMap(&smap)
	//fmt.Println(count)

	//userid := "test123"
	//sec := gauth.CreateSecret(gauth.GoogleAuthenticator)
	//url := gauth.CreateProvisionURIWithIssuer(userid, sec)
	//fmt.Println(`sec:`, sec)
	//fmt.Println(`url:`, url)
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
