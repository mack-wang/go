package main


// 可以把chan想象成1个管道

func producer(data chan int) { // 生产者(用data整型管道)
	for i := 0; i < 4; i++ {
		data <- i // 将数字1，2，3，4依次放入管道
	}
	close(data) // 关闭管道
}

func consumer(data chan int, done chan bool) { // 消费者(用data整型管道，done布尔管道)
	for x := range data { // data整型管道，接收从任意一个地方往data管道里放的数据。
		println("接收到：", x)
	}

	done <- true // 将true放入管道
}

func main() {
	done := make(chan bool)
	data := make(chan int)

	go producer(data)
	go consumer(data, done)

	<-done // 阻塞管道，直到消费者发回结束信号
}
