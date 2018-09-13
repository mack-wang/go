package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup        // 控制当前程序进程的开和关，避免goruntine还没执行完，进程就关了
	ready := make(chan struct{}) // 创建一个通道，进行阻塞

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done() // defer 会在全部执行完成后，才会触发，包括通道阻塞

			println(id, ":ready")
			<-ready // 阻塞，一直到接收到关闭的命令
			println(id, ":running")
		}(i)
	}

	time.Sleep(time.Second)
	println("Ready?Go!")
	close(ready) // 向ready发送关闭通道消息，但此时的作用是发送开始通知
	wg.Wait()    // 确保所有的wg.Done()完成，才关闭进程
}
