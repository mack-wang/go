package main

import "fmt"

// 可以看做堆栈，可以看做队列，可以看做任务池子
// 原则：先进先出
// 先put的，先被get到
// 被get后，将会被释放
// 超出容量的会被select执行default操作
type pool chan []byte

func newPool(cap int) pool {
	return make(chan []byte, cap)
}

func (p pool) get() []byte {
	var v []byte
	select {
	case v = <-p:
	default:
		v = make([]byte, 10)
	}
	return v
}

func (p pool) put(b []byte) {
	select {
	case p <- b:
	default:
	}
}

func main() {
	p := newPool(3)
	p.put([]byte{1, 2})
	p.put([]byte{2, 3})
	p.put([]byte{3, 4})
	fmt.Println(p.get())
	fmt.Println(p.get())
	p.put([]byte{3, 4})
	p.put([]byte{3, 5})
	fmt.Println(p.get())
	fmt.Println(p.get())
	fmt.Println(p.get())
}
