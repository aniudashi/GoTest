package goTest

/*
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。

将大整数加 1，并返回结果的数字数组。



示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
加 1 后得到 123 + 1 = 124。
因此，结果应该是 [1,2,4]。
示例 2：

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
加 1 后得到 4321 + 1 = 4322。
因此，结果应该是 [4,3,2,2]。
示例 3：

输入：digits = [9]
输出：[1,0]
解释：输入数组表示数字 9。
加 1 得到了 9 + 1 = 10。
因此，结果应该是 [1,0]。


提示：

1 <= digits.length <= 100
0 <= digits[i] <= 9
digits 不包含任何前导 0。

给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
考察：数组操作、进位处理--那需要数组直接操作
*/
func DigitArrayPlusOne(sliceInts []int) []int {
	intLength := len(sliceInts)
	for i := intLength - 1; i >= 0; i-- {
		intValue := sliceInts[i]
		if intValue == 9 {
			sliceInts[i] = 0 //最后一位是9，重置为0，往前处理一位
			continue
		} else {
			sliceInts[i]++
			return sliceInts //不是9，直接+1 返回
		}
	}
	// 走到这说明前面全部=9，结果是原数组长度+1，第一位是1后跟n个0
	retSlice := make([]int, len(sliceInts)+1)
	retSlice[0] = 1
	return retSlice
}
