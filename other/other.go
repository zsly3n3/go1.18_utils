package other

import (
	"fmt"
	"math/rand"
	"time"
)

// 洗牌算法
func Shuffle(arr []int) {
	rand.Seed(time.Now().UnixNano())
	var i, j int
	var temp int
	for i = len(arr) - 1; i > 0; i-- {
		j = rand.Intn(i + 1)
		temp = arr[i]
		arr[i] = arr[j]
		arr[j] = temp
	}
}

// 随机红包
// remainCount: 剩余红包数
// remainMoney: 剩余红包金额（单位：分)
func randomMoney(remainCount, remainMoney int) int {
	if remainCount == 1 {
		return remainMoney
	}
	rand.Seed(time.Now().UnixNano())
	var min = 1
	max := remainMoney / remainCount * 2
	money := rand.Intn(max) + min
	return money
}

// 发红包
// count: 红包数量
// money: 红包金额（单位：分)
func RedPackage(count, money int) {
	for i := 0; i < count; i++ {
		m := randomMoney(count-i, money)
		fmt.Printf("%d  ", m)
		money -= m
	}
}
