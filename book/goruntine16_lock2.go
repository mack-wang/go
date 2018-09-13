package main

// 将Mutex锁粒度控制在最小范围，及早释放

// 错误用法 1
func do() {
	//m.Lock()
	//url := cache["key"]
	//http.Get(url)  // 这个无须用锁，所以应该排除
	//m.Unlock
}

// 错误用法 2
func do2() {
	//m.Lock()
	//defer m.Unlock
	//url := cache["key"]
	//http.Get(url) // 同上
}

// 正确用法
func d3() {
	//m.Lock()
	//url := cache["key"]
	//m.Unlock
	//http.Get(url)
}

// 不支持递归锁
/*
m.Lock()
    m.Lock()
	m.Unlock()
m.Unlock()
*/

/*
注意：
1、对性能要求高时，应该避免使用defer Unlock
2、读写并发时，用RWMutex性能会更好
3、对单个数据读写时，尽量原子操作
4、要严格测试，尽量打开数据竞争检查

*/
