package main

func main() {
	// 普通用法
	// 如果i:=0中添加一个函数，这个函数只会执行1次
	for i := 0; i < 10; i++ {
		println(i)
	}

	// 相当于while i < 10,或者for ; x < 10 ; {}
	// 如果10换成一个函数，那么，这个函数每次循环都会执行
	i := 5
	for i < 10 {
		println(i)
		i++
	}

	// 相当于while true
	for {
		println(i)
		i++
		if i > 20 {
			break
		}
	}

}
