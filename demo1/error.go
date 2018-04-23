package main

import (
	"log"
	"fmt"
	"os"
)

func main()  {
	// 日志前缀
	log.SetPrefix("日志输出: ")
	// 日志前缀尾巴，可以是日期，时间等，012345
	log.SetFlags(0)

	log.Printf("你好像出错了,原因是%v","因为爱")

	fmt.Fprintf(os.Stderr, "Site is down: %v\n", "因为爱")
}
