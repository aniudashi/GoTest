// 包申明
package main

//引入
import (
	"fmt"
	"os"
)

func printString(s string) {
	fmt.Println(s)
}

var s1 string = "test2"

func main() {
	fmt.Println(os.Args)
	fmt.Println("测试go 程序")
	printString(s1)
}
