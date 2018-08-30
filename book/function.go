package main

import (
	"fmt"
)

// 函数的传入参数，返回参数均可以作为变量
// return未指明返回时，直接返回对应变量,隐式返回

// 如果返回值的类型能明确其含义，就尽量不要采用命名式的返回值
func say(name string) (age int) {
	println(name)
	age = 23
	return
}

func sing(words string) int{
	println(words)
	return len(words)
}

type Product struct {
	Name string
	Count int
	Size int
}

// 如果参数复杂，就应该将参数包装成复合结构类型
func shop(product Product)  {
	fmt.Println(product)
}


// 变参   实际上就是切片，只能放在尾部
func student(count ...int)  {
	for _,v := range count {
		println(v)
	}

	for i := range count {
		count[i] += 10
	}
}


// 匿名函数，好处是，可以在函数中嵌套写函数，函数还可以被返回，匿名函数还可以作为参数传入
func family() func(){
	i := 0
	return func(){
		i++
		println(i)
	}
}

// 匿名函数还可以写到结构体中
type Car struct {
	drive func(speed int)
}

func main() {
	age := say("tom")
	println(age)

	words := sing("hello")
	println(words)

	shop(Product{"a",2,2})

	data := []int{1,2,3,4}
	// 会改变data数据，如果不改变，需要操作复制数据
	student(data...)
	fmt.Println(data)
	data2 := []int{0,0,0,0}
	copy(data2,data)

	fmt.Println(data2)
	fmt.Println(data)

	// family() 仅返回匿名函数，并没有执行匿名函数
	family()()

	car := Car{func(speed int) {
		fmt.Println(speed)
	}}

	car.drive(200)

	hello := func() {
		println("hello 2")
	}

	hello()
}
