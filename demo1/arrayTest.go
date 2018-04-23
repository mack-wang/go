package main

import "fmt"

func main() {
	var a = [3]int{1, 2, 3};
	//首项
	fmt.Println(a[0])

	//尾项
	fmt.Println(a[len(a)-1])

	// 索引
	for index, item := range a {
		fmt.Println(index, item) // 前者是索引，后者是项
	}

	//　如果数组的长度未知，可以用省略号替换,但会根据赋值确定数组的长度
	b := [...]int{4, 5, 6, 7, 8}
	fmt.Println(len(b));

	//
	var c [5]int;
	//c = [5]int{1,2,3,4,5};
	c[0] = 1;
	c[1] = 2;
	for index, item := range c {
		fmt.Println(index, item) // 前者是索引，后者是项
	}

	// slice
	// []不添加数字或省略号 即为slice类型
	family := []string{0: "Dad", 1: "tom", 2: "bob", 3: "cindy"}
	for index, item := range family {
		fmt.Println(index, item) // 前者是索引，后者是项
	}

	family = append(family, "linda");
	for index, item := range family {
		fmt.Println(index, item) // 前者是索引，后者是项
	}

	family = append(family, "linda1", "linda2");
	for _, item := range family {
		fmt.Println(item)
	}
	fmt.Println("-------")
	family = append(family, family...);
	for index, item := range family {
		fmt.Println(index, item)
	}

}
