package main

func main() {
	var a [2]int
	var b [2]int
	//var c [3]int 数量不同，不是同类型，不可比较
	a[0] = 1
	a[1] = 2
	b[0] = 1
	b[1] = 3
	println(a == b)

	var e = [...]int{1,2,3:50}
	for _,v := range e {
		println(v)
	}
}