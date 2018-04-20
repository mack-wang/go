package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 变量声明规则
	// var 变量名字 类型 = 表达式

	// 默认值
	// int 0 / boolean false / string '' / slice、map、chan和函数 nil
	var num int
	var b bool
	var s string
	fmt.Println(num, b, s)

	// 同时赋值多个，同js
	var c, d, e = true, 2, "hello"
	fmt.Println(c, d, e)

	// 一组变量也可以通过调用一个函数，由函数返回的多个返回值初始化
	text, err := ioutil.ReadFile("./test.txt")

	// 简短变量声明 :=
	// :是声明 =是赋值
	// 简短变量声明更为常用，由赋值决定变量的类型
	// 而使用var声明的通常是要显示的表现出变量的类型

	// 如果之前已经声明过的变量，再使用简短变量声明，那么只保留赋值作用
	i := 100 // an int
	var boiling float64 = 100 // a float64
	var names []string
	var err error

	// 多个简短变量同时声明
	i, j := 0, 1
}
