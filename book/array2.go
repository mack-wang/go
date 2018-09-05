package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a
	println(&b, &a)
	b[2] = 9
	println(&b, &a)
	fmt.Println(b, a)

	b[0:3][2] = 8
	c := b
	println(&b, &c)
	fmt.Println(b, c, a)

	d := &a
	(*d)[1] = 7
	println(d)
	fmt.Println(b, c, a, *d)

	// 上面指针、切片等都会改变所有变量
}
