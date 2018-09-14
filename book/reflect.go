package main

import (
	"fmt"
	"reflect"
)

type X int

type user struct {
	name string
	age  int
}

type manager struct {
	user
	title string
}

type car struct {
	name string `field:"name" type:"varchar(50)"`
	age  int    `field:"age" type:"int"`
}

func (X) String() string {
	return ""
}

func main() {
	var a X = 100
	t := reflect.TypeOf(a)

	// 前者表示静态类型、后者表示基础类型（底层类型）
	fmt.Println(t.Name(), t.Kind()) // X int

	// Elem 能够返回指针、数组、切片、字典或通道的基类型
	var b = [3]int{1, 2, 3}
	fmt.Println(reflect.TypeOf(b))        // 返回 [3]int 复合类型
	fmt.Println(reflect.TypeOf(b).Elem()) // 返回 int 基类型

	var m manager
	d := reflect.TypeOf(m)
	name, _ := d.FieldByName("name") // 使用名称查找
	fmt.Println(name.Name, name.Type)

	age := d.FieldByIndex([]int{0, 1}) // 使用索引查找
	fmt.Println(age.Name, age.Type)

	// 可用反射提取struct tag , 还能自动 分解 常用于ORM映射或数据验证
	var ca car
	cc := reflect.TypeOf(ca)
	for i := 0; i < cc.NumField(); i++ {
		f := cc.Field(i)
		fmt.Println(f.Name, f.Tag.Get("field"), f.Tag.Get("type"))
	}

	// 高级用法
	var e X
	et := reflect.TypeOf(e)
	// 判断是否实现某接口，即X有没有实现String()接口
	ct := reflect.TypeOf((*fmt.Stringer)(nil)).Elem() // Stirng()接口
	fmt.Println(et.Implements(ct))                    // et是否实现了String()接口方法

	// 判断是否能将此类型，转换成其他类型，即X类型能否转成int类型
	rt := reflect.TypeOf(2)
	fmt.Println(et.ConvertibleTo(rt)) // 是否能转X转成bool

	// 判断是否能赋值给其他类型的变量,即其他类型的变量xxx = X, 能否赋值, X的基类型是int所以X=1没问题，但int类型的=X类型是不行的
	fmt.Println(et.AssignableTo(rt))

}
