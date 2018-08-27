package main

import "fmt"

func main() {
	// var const type 都可以合并成一个组，其实其实就是省去每次定义都要写相同的关键字
	type (
		user struct {
			Name string
			Age  int
		}
		event func(string) bool
	)

	type User struct {
		Name string
		Age  int
	}

	type Event func(string) bool

	userOne := user{
		"wlc",
		23,
	}
	fmt.Println(userOne)

	type data int
	var phone data = 100
	println(phone)
}
