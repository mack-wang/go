package main

func main() {
	exit := make(chan struct{})

	println("hello world") // 同步

	go func(s string) { // 并发
		println(s)
		close(exit)
	}("hello world 2")

	println("before exit") // 同步
	<-exit                 // 等待接收到退出信号，再继续执行
	println("exit")

}
