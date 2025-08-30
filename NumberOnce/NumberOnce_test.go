// 包申明
package goTest

import (
	"fmt"
	"testing"
)

//引入

func TestDoNumberOnce1(t *testing.T) {
	var sliceInt = []int{1, 2, 3, 4, 5, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceIntRet := DoNumberOnce(sliceInt)
	if len(sliceIntRet) == 1 && sliceIntRet[0] == 6 {
		fmt.Println("Test DoNumberOnce")
		t.Logf("输入%v DoNumberOnce 结果正常", sliceInt)
	}
}

func TestDoNumberOnce2(t *testing.T) {
	var sliceInt = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceIntRet := DoNumberOnce(sliceInt)
	if len(sliceIntRet) == 1 && sliceIntRet[0] == 6 { //if 变量需要写在前，常量在后 和c++习惯相反
		fmt.Println("Test DoNumberOnce")
		t.Logf("输入%v DoNumberOnce 结果正常", sliceInt)
	} else {
		t.Fatalf("输入%v DoNumberOnce 结果异常", sliceInt)
	}
}

//执行测试 最后需要跟原函数所在的go文件 go test  .\NumberOnce_test.go .\NumberOnce.go
//t.Fatalf 能正常输出 ，t.Logf 需要go test -v参数可以输出
// fmt和 t.Logf无输出-网上：设置.vscode 文件夹及配置 settings.json 文件中的 go.testFlags 参数为 '-v'（未实践）
