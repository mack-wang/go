package main

import "sync"

type receiver struct {
	sync.WaitGroup
	data chan int
}

// 工厂方法
func newReceiver() *receiver {
	r := &receiver{
		data: make(chan int),
	}

	r.Add(1)
	go func() {
		defer r.Done()
		for x := range r.data { // r.data会一直阻塞，一直读数据，直到被关闭，被关闭后r.Done才会被执行
			println("receive:", x)
		}
	}()

	return r
}

func main() {
	r := newReceiver()
	r.data <- 1
	r.data <- 2
	close(r.data)
	r.Wait()
}
