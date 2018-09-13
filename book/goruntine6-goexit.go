package main

import "runtime"

func main() {
	exit := make(chan struct{})

	go func() {
		defer close(exit)
		defer println("a")

		func() {
			defer func() {
				println("b", recover() == nil)
			}()

			func() {
				println("c")
				// 终止当前及以下的，类似并发中的break，不可以直接在非并发函数内调用这个，会崩溃
				runtime.Goexit()
				println("c done")
			}()

			println("b done")
		}()

		println("a done")
	}()

	<-exit
}
