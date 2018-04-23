package main

// 可变参数
// args ...interface{}可以接受任意参数类型
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main()  {

	println(sum(1,2,3,4))

	var a = []int{1,2,3,4,5}
	//println(a...)
	// slice可以作为参数，被展开
	println(sum(a...))
}
