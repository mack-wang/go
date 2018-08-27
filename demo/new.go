package main

import "fmt"

func main() {
	// p不是一个变量，是一个指针，是一个地址
	p := new(int)   // p, *int 类型, 指向匿名的 int 变量
	fmt.Println(*p) // "0"
	*p = 2          // 设置 int 匿名变量的值为 2
	fmt.Println(*p) // "2"
}
