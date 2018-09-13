package main

func main() {
	// 带有三个缓冲槽的异步通道
	// 通过debug可以看到c的结构，里面原本 *buf是*[0]int nil，后变成了*[3]int

	// 异步模式，在缓冲区未满，或数据还没读完前，不会阻塞
	c := make(chan int, 3)
	c <- 1 // 然后buf缓存数组[0]会保存1
	c <- 2
	c <- 3
	println(<-c) // 先进先出，会先从[0]中先取出1
	c <- 4
	println(<-c) // 虽然[0]的已经被取出，4放到了[0]的位置，但此时取出来的是[1]的2,还是先进先出的原则，4是后进的，后出。
	println(<-c)
	println(<-c)

	// 放的时候不能超过通道的容量
}
