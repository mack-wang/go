package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var map1 map[string]int
	fmt.Println(map1["a"]) // 能对map未赋值属性读，但不能写
	//map1["a"] = 1
	//fmt.Println(map1["a"])

	var map2 = make(map[string]int)
	println(map2["a"])
	map2["a"] = 1
	println(map2["a"])

	// 删除键值对
	map2["b"] = 2
	fmt.Println(map2)
	delete(map2, "b")
	fmt.Println(map2)

	// map不能并发读写，当已经有一个资源在读或写map，另一个资源不能同时读或写map，否则会报错
	// 所以可以使用读写锁

	var lock sync.RWMutex
	m := make(map[string]int)
	go func() {
		for {
			lock.Lock() // 先把m对象锁住，若有其同资源同时读写，则要排队
			m["a"] += 1
			lock.Unlock()
			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		for {
			lock.RLock()
			_ = m["b"]
			lock.RUnlock()
			time.Sleep(time.Microsecond)
		}
	}()

	select {}

	/*
		var m = make(map[string]int,1000)性能比下者要好，因为已经预先分配好了内存，不用动态分配内存
		var m = make(map[string]int)
	*/

}
