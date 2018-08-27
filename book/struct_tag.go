package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name     string `user name` //这引号里面的就是tag
	Password string `user password`
}

func main() {
	user := &User{"tom", "pass"}
	s := reflect.TypeOf(user).Elem() //通过反射获取type定义
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag) //将tag输出出来
	}
}
