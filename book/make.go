package main

import "fmt"

func main() {
	// 基本类型分为：普通类型、引用类型
	// 引用类型就是 slice、map、channel这三种类型,他们需要更复杂的存储结构
	// 引用类型初始化必须用make，以确保完成全部的内存分配和相关属性的初始化
	intSlice := make([]int, 0, 10)
	intSlice = append(intSlice, 10)
	intSlice = append(intSlice, 20)
	fmt.Println(intSlice)

	stringIntMap := make(map[string]int)
	stringIntMap["age"] = 12
	fmt.Println(stringIntMap)

}
