package main

import "fmt"

func main() {
	// 切片  切片本身是个只读对象，其工作机制类似数组指针的一种包装
	// 与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
	// 简单理解，切片就是把一个类数组元素[x:y]切割出来，构造出下面的slice
	// 切片结构
	/*
		type slice struct{
			array unsafe.Pointer
			len int	   // 表示切片当前可用元素的长度
			cap int    // 表示切片所引用的数组的真实长度，可理解为容量
		}
	*/

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(x[:])
	fmt.Println(x[2:5])
	fmt.Println(x[2:5:7])
	fmt.Println(x[4:])
	fmt.Println(x[:4])
	fmt.Println(x[:4:6])

	fmt.Println(cap(x[2:5:7]))

	// 先把x的3：9切出来，成为切片，再把切片中的索引2元素的值变为99
	// 因为切片里面是数组指针，所以元素指向的仍然是原数组，所以改变后，会影响原数组
	x[3:9][2] = 99
	x[2] = 99
	fmt.Println(x)

	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := a[:]
	c := a
	println(&a)
	println(&b)
	println(&c)
	fmt.Println(&b, b)
	fmt.Println(&c, c)

	// 直接初始化切片，使用make
	s1 := make([]int, 3, 5)
	fmt.Println(cap(s1))
	s1 = []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(s1)
	fmt.Println(cap(s1))

	// 使用数组时，无论什么操作，都要先复制整个数组元素，而使用将数组转成切片后，则无需复制整个数组，直接操作数组

	// 复制对象
	// 使用场景，如果小切片要长时间引用一个大数组，那可以使用复制的方法来新建切片，这样大数组内存就可以得到释放。
	ii := make([]int, 3)
	copy(ii, a)
	fmt.Println(ii)
}
