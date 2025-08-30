// 包申明
package goTest

//引入
import (
	"fmt"
)

/*
题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
考察：字符串处理、栈的使用
*/
func CheckValiadBraketsString(str string) bool {
	fmt.Println(str)

	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, ch := range str {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
			//push 加入最后
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
			// pop 去处最后一个
		}
	}
	return len(stack) == 0
}
