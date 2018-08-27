package main

func main() {
	const i, f = 0, 1
	println(i, f)

	const (
		x uint16 = 120
		y
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
		dd float64 = iota
		ee
	)
	println(aa, bb, cc, dd, ee)
}
