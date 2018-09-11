package main

func main() {
	// 添加go关键字就可以创建并发任务
	// 关键字go并非执行并发操作，而是创建了一个任务并发单元。新建任务被放到一个队列中，等待调度器去安排合适的线程执行。
	println("hello world")

	// 此时你发现，下面并没打印出来，进程退出并没有等并发任务执行完成
	go func(s string) {
		println(s)
	}("hello world 2")

}
