package main

import "fmt"

func main() {
	// 方法一：创建一个String=>int的hash map
	ages := make(map[string]int)

	ages["tom"] = 24;
	for index, i := range ages {
		// 此时的index变成了属性名
		fmt.Println(index, i)
	}

	// 方法二:
	ages2 := map[string]int{
		"bob":23,
		"tom":24,
	}

	delete(ages2, "tom")
	for index, i := range ages2 {
		// 此时的index变成了属性名
		fmt.Println(index, i)
	}

	// 注意map遍历时是随机的，不是按顺序，想要按顺序，得在之前先排序

	// 判断元素是否存在
	ages["linda"] = 25

	// ok用于判断是否存在
	if bobage,ok := ages2["bob"]; !ok {
		fmt.Println(bobage)
	}

}
