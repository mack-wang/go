package main

import "sync"

func main() {
	var wg sync.WaitGroup // 控制进程
	wg.Add(3)

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
			case x, ok = <-a:
				name = "a"
				if !ok {
					a = nil // 如果一个通道被设置为nil，他会被阻塞，并且不再被select选中
					break
				}
				println(name, x)
			case x, ok = <-b:
				name = "b"
				if !ok {
					b = nil
					break
				}
				println(name, x)
			}

			if a == nil && b == nil { // 如果有任意一个通道被关闭了，就停止，会触发defer
				return
			}

		}
	}()

	go func() {
		defer wg.Done()
		defer close(a)
		for i := 0; i < 3; i++ {
			a <- i
		}
	}()

	go func() {
		defer wg.Done()
		defer close(b) //
		for i := 0; i < 50; i++ {
			b <- i * 10
		}
	}()

	wg.Wait()

}
