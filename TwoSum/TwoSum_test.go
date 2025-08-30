package goTest

import (
	"reflect"
	"testing"
)

/*
示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]


提示：

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案


进阶：你可以想出一个时间复杂度小于 O(n2) 的算法吗？


*/

func TestDoTwoSum1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expectedNums := []int{0, 1}
	retNums := DoTwoSum(nums, target)

	if reflect.DeepEqual(expectedNums, retNums) {
		t.Logf("DoTwoSum 结果[%v]符合预期", retNums)
	} else {
		t.Fatalf("DoTwoSum 结果[%v]不符合预期", retNums)
	}
}

func TestDoTwoSum2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	expectedNums := []int{1, 2}
	retNums := DoTwoSum(nums, target)

	if reflect.DeepEqual(expectedNums, retNums) {
		t.Logf("DoTwoSum 结果[%v]符合预期", retNums)
	} else {
		t.Fatalf("DoTwoSum 结果[%v]不符合预期", retNums)
	}
}

func TestDoTwoSum3(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	expectedNums := []int{0, 1}
	retNums := DoTwoSum(nums, target)

	if reflect.DeepEqual(expectedNums, retNums) {
		t.Logf("DoTwoSum 结果[%v]符合预期", retNums)
	} else {
		t.Fatalf("DoTwoSum 结果[%v]不符合预期", retNums)
	}
}
