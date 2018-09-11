package main

import "fmt"

// 接口只能定义方法，不能有属性字段
// 接口不能自己去实现方法
type tester interface {
	say()
	pull() string
}

type data struct {
	name string
}

func (data) say() {
	fmt.Println("hello")
}

func (data) pull() string {
	return "hello"
}

// 接口不能自己去实现
//func (tester) say()()  {
//	println("hello")
//}

func main() {
	var a tester = data{}
	var i = []int{1, 2, 3}
	fmt.Println(i)
	i[0]++
	a.say()
	b := a.pull()
	println(b)
}
