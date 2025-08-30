package goTest

import (
	"fmt"
	"strings"
)

/*
1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 如果非空，则仅由小写英文字母组成

考察：字符串处理、循环嵌套
*/
func GetMaxCommonPrex(sliceStrs []string) (string, error) {
	strRet := ""
	if len(sliceStrs) <= 1 || len(sliceStrs) >= 200 {
		return "", fmt.Errorf("输入数组%v,长度不符合要求,长度=%d", sliceStrs, len(sliceStrs))
	}
	strCommon := sliceStrs[0]
	if len(strCommon) == 0 {
		return strRet, nil
	}
	sliceTmpStrs := sliceStrs[1:]
	strPrex := strCommon[0:1]
	for i := 1; i <= len(strCommon); i++ {
		bMatch := true
		for _, strValue := range sliceTmpStrs {
			if strings.HasPrefix(strValue, strPrex) {
				continue
			} else {
				bMatch = false
				break
			}
		}
		if bMatch {
			strRet = strPrex             //记录上一次满足条件的前缀
			strPrex = strCommon[0 : i+1] //全部都满足，前缀+1后再次开始判断
		} else {
			break
		}
	}
	return strRet, nil
}
