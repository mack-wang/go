package main

import "fmt"

func main() {
	// 如果常量在赋值时没有给定值，则会默认给定左侧的值
	const (
		a = 1
		b // 默认为1
		c = 2
		d // 默认为2
	)

	//iota 常量生成器
	// 一个特殊的符号，可以根据一定规则自动生成批量常量
	const (
		one = iota // 起始是0
		two
		three
		four
	)

	fmt.Println(a, b, c, d)            //1122
	fmt.Println(one, two, three, four) // 0123

	const (
		one1 = iota + 100 // 起始是0
		two1
		three1
		four1
	)
	fmt.Println(one1, two1, three1, four1) // 100 101 102 103

}
