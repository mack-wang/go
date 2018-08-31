package main

import "fmt"

func main() {
	var a [2]int
	var b [2]int
	//var c [3]int 数量不同，不是同类型，不可比较
	a[0] = 1
	a[1] = 2
	b[0] = 1
	b[1] = 3
	println(a == b)

	var e = [...]int{1, 2, 3: 50}
	for _, v := range e {
		println(v)
	}

	var d = [...]int{1, 2, 3, 4, 5}
	fmt.Println(d)

	var g = [][]int{
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println(g)

	// 多维数组，仅第1维可以加...
	var k = [...][]int{
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println(k)

	// 要想不影响原数组，就得用copy复制
	var l = []int{4, 5, 6}
	var j []int
	j = l
	j[2] = 7
	println(&l, &j)
	fmt.Println(l, j)

	var m *[]int
	m = &l
	// *[]int是不支持索引操作的，所以得先返回[]int后再索引
	(*m)[2] = 8
	j[2] = 9
	fmt.Println(l, j, *m)

	ll := []int{1, 2, 3, 4, 100}
	// 函数传参
	func1 := func(data []int) {
		fmt.Println(data)
	}

	func1(ll)
	func1(ll[:])

}
