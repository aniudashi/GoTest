// 包申明
package goTest

import "fmt"

func DoNumberOnce(sliceInt []int) []int {
	var mpInt map[int]int = make(map[int]int)
	for _, valInt := range sliceInt {
		if _, exist := mpInt[valInt]; exist {
			var cnt = mpInt[valInt]
			mpInt[valInt] = cnt + 1
		} else {
			mpInt[valInt] = 1
		}
	}
	var sliceIntRet = make([]int, 0, 0) //是不是可以考虑优化
	for index, valInt := range mpInt {
		if valInt == 1 {
			sliceIntRet = append(sliceIntRet, index)
		}
	}
	fmt.Println("sliceIntRet=", sliceIntRet)
	return sliceIntRet
}
