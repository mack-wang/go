package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	c := make(chan int)
	var send chan<- int = c
	var recv <-chan int = c

	// 只接收
	go func() {
		defer wg.Done()

		for x := range recv {
			println(x)
		}
	}()

	// 只发送
	go func() {
		defer wg.Done()
		defer close(c)

		for i := 0; i < 3; i++ {
			send <- i
		}
	}()

	/*
		不可以做逆向操作
		<-send
		recv <- i
	*/

	/*
		不可以close 接收端 recv
	*/

	/*
		单向通道无法转换回去
	*/

	wg.Wait()
}
