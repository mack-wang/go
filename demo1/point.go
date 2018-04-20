package main

import "fmt"

func main()  {
	x := 1
	p := &x // p是一个指针，*int类型 是一个32位的int类型的16进制数字，是一个地址，指向x变量的具体内容
	fmt.Println(p) // "1"
	fmt.Println(*p) // "1"
	*p = 2 // equivalent to x = 2
	fmt.Println(x) // "2"
}
