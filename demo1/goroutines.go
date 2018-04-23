package main

import (
	"fmt"
	"time"
)

func main() {
	// go 开通另一线程来运行
	// 继续执行下面的动作
	//　如果不加go，则会停止5秒才继续执行
	go time.Sleep(5 * time.Second)
	Go()
}
func Go() {
	fmt.Println("Go Go Go!!!")
}