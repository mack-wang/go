package main

import "time"


//　只有首字母大写的属性，才会被导出（即包外可见），
// 首字母大写，相当于public, 小写相当于private
type Employee struct {
	ID int
	Name string
	Address string
	DoB time.Time
	Position string
	Salary int
	ManagerID int
}
var dilbert Employee

type Color struct {
	red int
	blue int
	green int
}

type Car struct {
	CarCan
	name string
	price int
}

type CarCan struct {
	speed int
}


func main()  {
	dilbert.Salary = 5000
	println(dilbert.Salary)

	color := Color{225,255, 305}
	println(color.red, color.blue, color.green)

	// 如果考虑效率的话，较大的结构体通常会用指针的方式传入和返回
	var car Car
	car.name = "benz"
	car.price = 35
	car.speed = 120
	println(car.name,car.price,car.speed)
}
