package main

import "fmt"

// 定义手机接口
type Phone interface {
	call()
}

// 定义类型
type NokiaPhone struct {
}

// 使用NokiaPhone类型，实现接口方法
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

// 使用IPhone类型，实现接口方法
func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	// 创建Phone类型，继承了call方法
	var phone Phone

	// 根据不同的手机类型实现了不同的call的具体实现方法
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}
