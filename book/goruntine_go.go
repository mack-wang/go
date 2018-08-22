package main

import (
	"fmt"
	"time"
)

func main() {
	// 5个并发执行，所以每秒会同时输出5个结果
	go task(1)
	go task(2)
	go task(3)
	go task(4)
	go task(5)

	time.Sleep(time.Second * 6)
}

func task(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d: %d\n", id, i)
		time.Sleep(time.Second)
	}
}
