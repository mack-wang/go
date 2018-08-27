package main

import "fmt"

func main() {
	// 经典for循环
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 省略初始化
	var i = 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	// 省略全部，无限循环，break跳出
	for {
		var a = 100
		fmt.Println(a)
		break
	}

	// ++ -- 自增 自减
	var b = 100
	for {
		b++
		if b == 105 {
			break
		}
		fmt.Println(b)
	}
}
