package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 只有所有属性都可比较时，struct结构体对象才能比较
	type Data struct {
		age int
	}

	a := Data{1}
	b := Data{2}
	c := Data{1}

	fmt.Println(a == b)
	fmt.Println(a == c)

	// 所有字段必须唯一，但可以用“_”进行补位,并且"_"可以重复
	type HomeData struct {
		_     int
		age   int
		gay   int
		my    int
		_     string
		hello int
	}

	d := HomeData{1, 2, 3, 4, "6", 9}
	fmt.Println(d)

	// 可以使用指针操作对象，但不能使用多级指针
	e := &d
	(*e).age = 3
	fmt.Println(e)

	// 匿名字段
	type family struct {
		int // 会默认使用类型作为名字，效果同 int int
		age int
	}

	f := family{1, 2}
	fmt.Println(f)
	fmt.Println(f.int)
	fmt.Println(f.age)

	// go不是传统的面向对象的编程语言，仅实现了最小的面向对象机制

	// tag有点像java中的注解，虽然像注释，但具有一定的意义，通过反射可以获取。可以为属性带来额外的可配置性
	type User struct {
		Name string `info`
		Age  int    `45`
	}

	u := User{"tom", 23}

	fmt.Println(reflect.TypeOf(u).Field(0).Name)
	fmt.Println(reflect.TypeOf(u).Field(0).Tag)
	fmt.Println(reflect.TypeOf(u).Field(0).Index)

}
