package main



func main() {

	println(1)

	println(2)

	i := 0

	three:
	println(3)

	if i < 2 {
		i++
		// goto能跳转到标签处再继续执行
		// 注意跳出条件，不然就会一直循环

		// goto只能跳转到同一函数作用域内
		goto three
	}


}
