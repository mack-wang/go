package main

func main() {
	a, b, c, x := 5, 4, 3, 4
	x--

	// 相当于多个if
	// 默认每个case都有break,也可以显示的写break
	switch {
	case x == a:
		println(x)
	case x == b:
		println(-x)
	case x == c:
		println(c)
	default:
		println(0)
	}

	// x不会作为switch传参传入，但会执行这个计算
	switch x++; {
	case x == a:
		println(x)
	case x == b:
		println(-x)
	case x == c:
		println(c)
	default:
		println(0)
	}

	// 会将x与case多个结果比对
	x++
	switch x {
	case a:
		println(x)
	case b:
		println(-x)
	case c:
		println(c)
	default:
		println(0)
	}

	// 会先执行计算
	switch x = 3; x {
	case a:
		println(x)
	case b:
		println(-x)
	case c:
		println(c)
	default:
		println(0)
	}

	// 如果想实现没有break继续执行的效果，要加fallthrough
	switch x = 3; x {
	case a:
		println(x)
		fallthrough
	case b:
		println(-x)
	case c:
		println(c)
	default:
		println(0)
	}
}
