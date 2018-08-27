package main

import "fmt"

func main() {
	fmt.Println(len("hello world"))
	// 一个中文占用3个字节
	fmt.Println(len("你好"))

	// 会截取0-5的字节字符，且不包括5
	fmt.Println("hello world"[0:5])

	// 你
	fmt.Println("你好"[0:3])
	// 你? 因为第4个字节无法构成好字，3个字节才能
	fmt.Println("你好"[0:4])

	// 字符串是不可改变的,要想改变，只能重新赋值
	s := "hello"
	fmt.Println(s)
	//s[0] = "a";
	s += " world"
	fmt.Println(s)
}
