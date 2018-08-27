package main

func main() {
	const i, f = 0, 1
	println(i, f)

	const (
		x uint16 = 120
		y        // 会和上一变量相同
		s = "abc"
		z
	)
	println(x, y, s, z)
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
	)
	println(KB, MB, GB)
	const (
		a = iota
		b
		c = 100
		d
		e = iota // 4,会继续
		g
	)
	println(a, b, c, d, e, g)
	const (
		aa = iota
		bb
		cc
		dd float64 = iota // 会改变类型，但iota依然会继续
		ee
	)
	println(aa, bb, cc, dd, ee)
}
