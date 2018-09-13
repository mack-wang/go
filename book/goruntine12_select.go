package main

import "sync"

func main() {
	var wg sync.WaitGroup // 控制进程
	wg.Add(2)

	a, b := make(chan int), make(chan int) // 创建2个通道

	go func() {
		defer wg.Done() // 确保两个goruntine都会被执行完成

		for {
			var (
				name string
				x    int
				ok   bool
			)

			// 随机选择通道读取数据
			select {
			case x, ok = <-a: // 即使这两个case都是a，那也会随机选择
				name = "a"
			case x, ok = <-b:
				name = "b"
			default: // 记得加default，避免select阻塞

			}

			if !ok { // 如果有任意一个通道被关闭了，就停止，会触发defer
				return
			}

			println(name, x)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(a)
		defer close(b) //
		for i := 0; i < 10; i++ {
			select { // 随机选择通道发送消息，a,b发送数据次数也是随机的
			case a <- i:
			case b <- i * 10:
			}

		}
	}()

	wg.Wait()

}
