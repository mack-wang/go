package main

import (
	"runtime"
)

func main() {
	// Go从1.5版本开始，默认采用多核执行，默认是你的CPU核心数，以前版本默认为1
	println(runtime.NumCPU())
	runtime.GOMAXPROCS(8)
}
