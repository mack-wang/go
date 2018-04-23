package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //设置使用的CPU核数（获取本台计算机的CPU核数）
	// 创建channel 管道
	c := make(chan bool)
	// 创建10条管道
	for i := 0; i < 10; i++ {
		go Go1(c, i)
	} //当一个核的时候还是按部就班的一个线程一个线程开始执行，可是现在变成多核以后，指不定先执行哪个后执行哪个，可能先执行i=5的，也可能先执行i=8的，即分配任务不定的特性

	// 接收c的结果，并丢弃这个结果true，同时停止了channel
	<-c
}
func Go1(c chan bool, index int) {
	a := 1
	// 进行1亿次加法
	for i := 0; i < 100000000; i++ {
		a += i
	}
	// 打印出
	fmt.Println(index, a)

	// 当index = 9时，传递true给c
	if index == 9 {
		c <- true //因此这个地方就错误了，有时候index=9并不是最后一个执行的。
	}
}
//输出的结果是不确定的，因为执行index=9的次序是不确定的。