package main

type Age int

// 如果方法内部不引用到实例，可以不传参数名，只传类型
// 最好传递一个有意义的名字，不要用this, self
func (age Age) say() {
	println(age)
}

func (age *Age) sayHello() {
	println(*age)
}

func main() {
	// 方法和函数的区别：方法是和实例绑定的，方法	属性某个类型，所以只有某个类型能调用，在go中会传递方法宿主类型，在js中会传递this
	var a Age = 3
	a.say()
	a = 4
	// TODO:p132
	a.sayHello()
}
