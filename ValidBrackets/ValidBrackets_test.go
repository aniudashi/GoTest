// 包申明
package goTest

import (
	"testing"
)

func TestCheckValiadBraketsString1(t *testing.T) {
	strTest1 := "()"
	if CheckValiadBraketsString(strTest1) == true {
		t.Logf("输入%s CheckValiadBraketsString 结果正常", strTest1)
	} else {
		t.Fatalf("输入%s CheckValiadBraketsString 结果异常", strTest1)
	}
}

func TestCheckValiadBraketsString2(t *testing.T) {
	strTest1 := "()[]{}"
	if CheckValiadBraketsString(strTest1) == true {
		t.Logf("输入%s CheckValiadBraketsString 结果正常", strTest1)
	} else {
		t.Fatalf("输入%s CheckValiadBraketsString 结果异常", strTest1)
	}
}

func TestCheckValiadBraketsString3(t *testing.T) {
	strTest1 := "(]"
	if CheckValiadBraketsString(strTest1) == false {
		t.Logf("输入%s CheckValiadBraketsString 结果正常", strTest1)
	} else {
		t.Fatalf("输入%s CheckValiadBraketsString 结果异常", strTest1)
	}
}

func TestCheckValiadBraketsString4(t *testing.T) {
	strTest1 := "([])"
	if CheckValiadBraketsString(strTest1) == true {
		t.Logf("输入%s CheckValiadBraketsString 结果正常", strTest1)
	} else {
		t.Fatalf("输入%s CheckValiadBraketsString 结果异常", strTest1)
	}
}

func TestCheckValiadBraketsString5(t *testing.T) {
	strTest1 := "([)]"
	if CheckValiadBraketsString(strTest1) == false {
		t.Logf("输入%s CheckValiadBraketsString 结果正常", strTest1)
	} else {
		t.Fatalf("输入%s CheckValiadBraketsString 结果异常", strTest1)
	}
}

func TestCheckValiadBraketsString6(t *testing.T) {
	strTest1 := "" //给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效,---不存在空的情况
	if CheckValiadBraketsString(strTest1) == false {
		t.Logf("输入%s CheckValiadBraketsString 结果正常", strTest1)
	} else {
		t.Fatalf("输入%s CheckValiadBraketsString 结果异常", strTest1)
	}
}
