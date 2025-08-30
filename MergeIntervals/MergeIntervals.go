package goTest

import (
	"sort"
)

/*
以数组 intervals
表示若干个区间的集合，
其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，
并返回一个不重叠的区间数组，
该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，
然后使用一个切片来存储合并后的区间，
遍历排序后的区间数组，
将当前区间与切片中最后一个区间进行比较，
如果有重叠，则合并区间；
如果没有重叠，则将当前区间添加到切片中。
*/

type Pair struct {
	left  int
	right int
}

type pair Pair
type pairSlice []pair

// 返回切片的长度
func (p pairSlice) Len() int {
	return len(p)
}

// 比较方法
func (p pairSlice) Less(i, j int) bool {
	return p[i].left < p[j].left
}

// 交换方法
func (p pairSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//完成sort.interface三个方法， len长度 swap交换 less比较规则，就可以通过sort排序

func DoMergeIntervals(slicePairs pairSlice) pairSlice {
	sort.Sort(pairSlice(slicePairs))           //排序
	pairSliceMerge := pairSlice{slicePairs[0]} //第一个区间设置为初始化数据
	for i := 1; i < len(slicePairs); i++ {     //依次往后比
		pairValueLast := &pairSliceMerge[len(pairSliceMerge)-1] //这里 需要为指针，不然修改的是副本
		if slicePairs[i].left <= pairValueLast.right {          //当前节点开始小于前一个的结束
			if slicePairs[i].right > pairValueLast.right { //当前节点结束大于前一个结束，设置前一个结束为当前的结束值
				pairValueLast.right = slicePairs[i].right //重叠合并
			} //else为当前节点前开始结束都被前一个包围，不需要处理
		} else { //不重叠
			pairSliceMerge = append(pairSliceMerge, slicePairs[i])
		}

	}

	return pairSliceMerge
}
