// 包申明
package goTest

import (
	"testing"
)

func TestCheckIsPalindrome1(t *testing.T) {
	intTest1 := 1234321
	if CheckIsPalindrome(intTest1) == true {
		t.Logf("输入%d CheckIsPalindrome 结果正常", intTest1)
	} else {
		t.Fatalf("输入%d CheckIsPalindrome 结果异常", intTest1)
	}
}

func TestCheckIsPalindrome2(t *testing.T) {
	intTest1 := 12343211
	if CheckIsPalindrome(intTest1) == true {
		t.Logf("输入%d CheckIsPalindrome 结果正常", intTest1)
	} else {
		t.Fatalf("输入%d CheckIsPalindrome 结果异常", intTest1)
	}
}
