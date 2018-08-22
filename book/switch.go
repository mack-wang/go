package main

func main(){
	x := -1
	switch {
	case x > 0:
		println(x)
	case x < 0:
		println(-x)
	default:
		println(0)
	}
}
