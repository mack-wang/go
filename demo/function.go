package main

import "fmt"

/*


func name(parameter-list 参数列表) (result-list 返回值列表) {
	body
}

*/

func say(name string, age int) (name1 string, age1 int) {
	name += "hello"
	age++
	return name, age
}

// 返回一个匿名函数，并执行
// Go使用闭包（closures）技术实
// 现函数值，Go程序员也把函数值叫做闭包。
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	name, age := say("bob", 23)
	println(name, age)

	// 函数可以被当作值，被赋值给其他变量，同js
	var s = say
	name2, age2 := s("tom", 24)
	println(name2, age2)

	// 通过这个例子，我们看到变量的生命周期不由它的作用域决定：squares返回后，变量x仍然
	//隐式的存在于f中。
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
}
