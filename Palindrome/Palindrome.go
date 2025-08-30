// 包申明
package goTest

//引入
import (
	"fmt"
)

/*
考察：数字操作、条件判断
题目：判断一个整数是否是回文数
*/
func CheckIsPalindrome(num int) bool {
	fmt.Println(num)
	if num < 0 {
		return false
	}
	var oriNumber int = num
	var reverseNumber int = 0
	for num > 0 {
		reverseNumber = reverseNumber*10 + num%10
		num = int(num / 10)
	}
	return oriNumber == reverseNumber
}
