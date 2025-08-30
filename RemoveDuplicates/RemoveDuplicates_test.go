package goTest

import (
	"testing"
)

/*
系统会用下面的代码来测试你的题解:

int[] nums = [...]; // 输入数组
int[] expectedNums = [...]; // 长度正确的期望答案

int k = removeDuplicates(nums); // 调用

assert k == expectedNums.length;

	for (int i = 0; i < k; i++) {
	    assert nums[i] == expectedNums[i];
	}

如果所有断言都通过，那么您的题解将被 通过。

示例 1：

输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
示例 2：

输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。
*/

func TestDoRemoveDuplicates1(t *testing.T) {
	nums := []int{1, 1, 2}      // 输入数组
	expectedNums := []int{1, 2} // 长度正确的期望答案

	k := DoRemoveDuplicates(nums) // 调用

	if k == len(expectedNums) {
		t.Logf("removeDuplicates 结果[%d]符合预期", k)
	} else {
		t.Fatalf("removeDuplicates 结果[%d]不符合预期", k)
	}
	for i := 0; i < k; i++ {
		if nums[i] == expectedNums[i] {
			t.Logf("removeDuplicates 结果[%d][%d]符合预期", i, nums[i])
		} else {
			t.Fatalf("removeDuplicates 结果[%d][%d]不符合预期", i, nums[i])
		}
	}

}

func TestDoRemoveDuplicates2(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4} // 输入数组
	expectedNums := []int{0, 1, 2, 3, 4}        // 长度正确的期望答案

	k := DoRemoveDuplicates(nums) // 调用

	if k == len(expectedNums) {
		t.Logf("removeDuplicates 结果[%d]符合预期", k)
	} else {
		t.Fatalf("removeDuplicates 结果[%d]不符合预期", k)
	}
	for i := 0; i < k; i++ {
		if nums[i] == expectedNums[i] {
			t.Logf("removeDuplicates 结果[%d][%d]符合预期", i, nums[i])
		} else {
			t.Fatalf("removeDuplicates 结果[%d][%d]不符合预期", i, nums[i])
		}
	}

}
