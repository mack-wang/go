package main

import "fmt"
import (
	"utils"
)

func main() {
	fmt.Println(utils.CToF(80))

	a := utils.Person{"bob", 23}
	b := utils.Person{"bob", 24}
	println(utils.Sum(a, b))
	println(a.Sum(b))
	println(a.Sum2(b))
	println(b.Mine(a))

	// 小写的方法默认private
	//println(b.mine(a))
}
