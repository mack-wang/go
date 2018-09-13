package main

func main() {
	done := make(chan struct{})
	done2 := make(chan struct{})
	msg := make(chan string)
	msg2 := make(chan string)

	// 同步模式，要配合goruntine操作
	go func() {
		s := <-msg
		println(s)
		msg2 <- "hello"
		close(done)
	}()

	go func() {
		s := <-msg2
		println(s)
		close(done2)
	}()

	msg <- "hi" // msg 通道造好；s:=<-msg阻塞中；msg<-"hi"往通道中传入消息; s:=<-msg接收到消息，解除阻塞，并打印;
	<-done      // 阻塞，直到有数据或管道关闭
	<-done2
}
