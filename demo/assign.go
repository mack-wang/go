package main

import "fmt"

func main() {
	s := "t"            // 最简洁，但只能在函数内部用
	var a string        // 声明类型
	var b = "t1"        // 同js
	var c string = "t2" // 声明并赋值
	fmt.Println(s, a, b, c)

	// 元组赋值

	aa, bb, cc := c, b, a
	fmt.Println(aa, bb, cc)
}
