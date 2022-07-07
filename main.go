package main

import (
	"fmt"
	"github.com/zsly3n3/go1.18_utils/sortBy"
	"golang.org/x/exp/constraints"
)

func main() {
	s1 := make([]float32, 0)
	s1 = append(s1, 5, 3, 2)
	s1 = append(s1, 1.0, 0, 9)
	s1 = append(s1, 2.5, 8, 11)
	num := len(s1)
	sortBy.NumericMergeSort(false, s1, 0, num-1)
	Test(s1)
}

func Test[T constraints.Ordered](arr []T) {
	for _, v := range arr {
		fmt.Println(v)
	}
}
