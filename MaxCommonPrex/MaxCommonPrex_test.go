package goTest

import "testing"

func TestGetMaxCommonPrex1(t *testing.T) {
	var strsInput = []string{}
	strRet, err := GetMaxCommonPrex(strsInput)
	if err != nil {
		t.Fatalf("输入%v GetMaxCommonPrex 结果[%s]不符合预期,错误%s", strsInput, strRet, err)
	}
	if strRet == "" {
		t.Logf("输入%v GetMaxCommonPrex 结果[%s]符合预期", strsInput, strRet)
	} else {
		t.Fatalf("输入%v GetMaxCommonPrex 结果[%s]不符合预期", strsInput, strRet)
	}

}

func TestGetMaxCommonPrex2(t *testing.T) {
	var strsInput = []string{"flower", "flow", "flight"}
	strRet, err := GetMaxCommonPrex(strsInput)
	if err != nil {
		t.Fatalf("输入%v GetMaxCommonPrex 结果[%s]不符合预期,错误%s", strsInput, strRet, err)
	}
	if strRet == "fl" {
		t.Logf("输入%v GetMaxCommonPrex 结果[%s]符合预期", strsInput, strRet)
	} else {
		t.Fatalf("输入%v GetMaxCommonPrex 结果[%s]不符合预期", strsInput, strRet)
	}

}

func TestGetMaxCommonPrex3(t *testing.T) {
	var strsInput = []string{"dog", "racecar", "car"}
	strRet, err := GetMaxCommonPrex(strsInput)
	if err != nil {
		t.Fatalf("输入%v GetMaxCommonPrex 结果[%s]不符合预期,错误%s", strsInput, strRet, err)
	}
	if strRet == "" {
		t.Logf("输入%v GetMaxCommonPrex 结果[%s]符合预期", strsInput, strRet)
	} else {
		t.Fatalf("输入%v GetMaxCommonPrex 结果[%s]不符合预期", strsInput, strRet)
	}

}
