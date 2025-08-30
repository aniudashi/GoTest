package goTest

import (
	"reflect"
	"testing"
)

/*
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
*/
func TestDigitArrayPlusOne1(t *testing.T) {
	digits := []int{1, 2, 3}
	digits = DigitArrayPlusOne(digits)
	//if digits == []int{1, 2, 4} //不支持直接==
	if reflect.DeepEqual(digits, []int{1, 2, 4}) {
		t.Logf("DigitArrayPlusOne 结果[%v]符合预期", digits)
	} else {
		t.Fatalf("DigitArrayPlusOne 结果[%v]不符合预期", digits)
	}

}

func TestDigitArrayPlusOne2(t *testing.T) {
	digits := []int{4, 3, 2, 1}
	digits = DigitArrayPlusOne(digits)
	//if digits == []int{1, 2, 4} //不支持直接==
	if reflect.DeepEqual(digits, []int{4, 3, 2, 2}) {
		t.Logf("DigitArrayPlusOne 结果[%v]符合预期", digits)
	} else {
		t.Fatalf("DigitArrayPlusOne 结果[%v]不符合预期", digits)
	}

}

func TestDigitArrayPlusOne3(t *testing.T) {
	digits := []int{9}
	digits = DigitArrayPlusOne(digits)
	//if digits == []int{1, 2, 4} //不支持直接==
	if reflect.DeepEqual(digits, []int{1, 0}) {
		t.Logf("DigitArrayPlusOne 结果[%v]符合预期", digits)
	} else {
		t.Fatalf("DigitArrayPlusOne 结果[%v]不符合预期", digits)
	}

}
