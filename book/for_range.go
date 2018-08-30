package main

import (
	"fmt"
	"time"
)

func main() {
	// for range 支持字符串、数组、数组指针、切片、字典、通道类型、

	// 返回索引和键值
	data := []string{"a","b","c"}
	for i, v := range data {
		println(i, v)
	}


	var i = 0
	for range data {
		// 可以不赋值，仅遍历
		i++
		println(i)
	}

	time.Sleep(1000)

	intData := []int{1}
	//for i := range intData {
	//	intData[i] += 1
	//	fmt.Println("a",intData)
	//}

	for i := range intData[:] {
		intData[i] += 1
		fmt.Println("b",intData)
	}

	time.Sleep(1000)

	for range intData {
		fmt.Println("c",intData)
	}


}
