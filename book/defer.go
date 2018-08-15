package main

func main()  {
	// 常用来释放资源，解除锁定，或执行一些清理操作
	// 可以定义多个defer，越早出现的，越晚执行
	defer println("结束后清理1...")
	defer println("结束后清理2...")
	defer println("结束后清理3...")

	println(10/1)
}