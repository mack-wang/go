package main

import "fmt"


// 除非是不可恢复性的，会导致系统异常的，导致无法正常工作的，否则不建议使用panic
func main() {
	defer func(){
		if catch := recover(); catch != nil{
			fmt.Println(catch)
		}
	}()

	// panic相当于throw 抛出一个异常， 会终止程序，并立即调用defer函数， recover会执行panic抛出的异常
	panic("error")
	panic("error2")
	panic("error3")
	println("exit.")
}
