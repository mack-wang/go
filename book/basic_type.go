package main

import (
	"fmt"
	"log"
)

func main() {
	var a byte = 1
	var b uint8 = 1
	//var c int = 1
	fmt.Println(a == b)
	//fmt.Println(a==c) // 错误

	var d float32
	fmt.Printf("%e\n", d)
	log.Println(d)

	fmt.Println(&a)
	fmt.Printf("%d\n", &a)

	// string 默认为空字符串，非nil
	var e string
	fmt.Println(e)
	fmt.Println(e == "")

	// struct function interface map slice channel 初始化全部为nil
}
