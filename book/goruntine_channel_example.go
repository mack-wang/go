package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second) // working
	fmt.Println("done")
	done <- true
}

// 只发送数据
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 只接收数据
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	// 一、最简单的channel使用
	// 1、创建channel管道 2、函数中往管道中放数据，并执行函数 3、从管道中取数据
	messages := make(chan string)
	go func() { messages <- "ping" }()
	msg := <-messages
	println(msg)

	// 实践,从数据库中获取用户账号名，并输出
	username := make(chan string)
	go func() {
		name := "tom"
		username <- name
	}()
	println(<-username)

	// 二、用于channel的range是阻塞的
	queue := make(chan string,2) // 不带数字2，分配固定数量缓存，则会陷入死锁
	queue <- "one"
	queue <- "two"
	close(queue)
	for elem := range queue {
		println(elem)
	}

	// 三、通道同步
	done := make(chan bool, 1)
	go worker(done)
	<-done //blocking 阻塞在这里,知道worker执行完毕

	// 四、限制通道的发送方向（只发送数据通道，只接收数据）
	pings := make(chan string,1)
	pongs := make(chan string,1)
	// 只发送数据
	ping(pings, "passed message")
	// 只接收数据
	pong(pings, pongs)
	// 打印接收到的数据
	fmt.Println(<-pongs)

}
