package main

import "runtime"

func main() {
	exit := make(chan struct{})

	// b 放到 a内部，是为了让a的优先级更高，更先执行
	// runtime.Gosched() 会让出当前线程，执行调试b
	// 但从测试结果来看，让出线程虽然大部分成功，但不是100%成功
	go func() {
		defer close(exit)

		go func() {
			println("b")
		}()

		for i := 0; i < 4; i++ {
			println("a", i)
			if i == 1 {
				runtime.Gosched()
			}
		}
	}()

	<-exit
}
