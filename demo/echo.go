package main

import (
	"fmt" //相当于print
	"os"  //系统，命令行参数os.Args
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
