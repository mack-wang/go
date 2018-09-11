package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// 当要等待多个并发任务执行完毕的时候，用sync.WaitGroup
	// wg.Add(1) 会给state1[1]，增加1个计数，go中defer wg.Done() 会减少一个计数，wg.Wait会一直阻塞，直到wg的state1[1]为0，即并发全部执行完毕，才会继续执行
	// defer wg.Done() 是为了确保先执行了sleep和println后，才执行wg.Done()
	// 如果没有sleep的话，可以不用defer，并把wg.Done()放到最后收尾
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			time.Sleep(time.Second)
			println("goroutine:", id, "done.")
		}(i)
	}

	println("main...")
	wg.Wait()
	println("main exit.")

	// 可同步在多处使用wg.Wait()，比如在go func中

}
