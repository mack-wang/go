package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 实例化input
	input := bufio.NewScanner(os.Stdin)
	// 调用input的scan 方法读取输入
	// 无限循环，一直读取
	for input.Scan() {
		// 输出读取的文本
		fmt.Println(input.Text())
	}
}
