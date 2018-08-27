package main

import "strconv"

func main() {
	if true {
		println(true)
	}
	x := 1
	if x++; x == 2 {
		println("先计算x++，再判断")
	}
	if c := 3; x == 2 {
		println("先赋值" + strconv.Itoa(c))
	}
}
