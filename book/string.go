package main

import (
	"fmt"
	"strings"
)

func main() {
	// `` 模板字符串，不做任何转义处理
	println(`hello
   ddd
hello`)

	// "" 普通字符串，做转义处理
	println("hello\nhellohello")

	// 字符串是不可变字符（byte）序列

	s:= "哈喽\x61\142"
	fmt.Printf("%s\n",s)
	fmt.Printf("% x\n",s)

	// 修改字符串，都必须先将字符串转成[]byte或[]rune类型才可以,然后还要转回来
	d:="hello"
	fmt.Println([]byte(d))
	c := []byte(d)
	c[2] = byte(66)
	c[3] = byte('c')
	println(string(c))

	println(strings.Repeat("tom", 10))

	f := []byte("lll")
	f = append(f, "eee"...)
	println(string(f))
}