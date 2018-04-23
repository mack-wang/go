package main

import "fmt"

//func main() {
//	c := make(chan bool) //make(chan [所存储的数据类型])
//	go func() {
//		fmt.Println("Go Go Go!!!")
//		c <- true //把true给到c
//	}()
//	<-c //把c取出来的操作，在这里会等待并发函数把数据给到c之后才执行，保证了并发线程结束后主线程才结束
//}


func main() {
	c := make(chan []int) //make(chan [所存储的数据类型])，直接make的都是双向通道，又可以存又可以取
	go func() {
		var d []int
		for i:=0;i<10;i++{
			d = append(d, i*26)
			c <- d //把d给到c
		}
		close(c)
	}()
	fmt.Println("---")
	// v 和 这个c 相等，每次c被重新赋值时，都会触发这个v = c
	for v := range c {
		fmt.Println(len(v))
		fmt.Println(v)
	}//一定要注意在某处把channel关闭掉，否则会成为死锁，无限等待
	fmt.Println("---2")
}