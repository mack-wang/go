package utils

type Person struct {
	Name string
	Age int
}

func Sum(a Person, b Person) int {
	return a.Age + b.Age
}

// a相当于this
func (a Person) Sum(b Person) int  {
	return a.Age + b.Age
}

// a相当于this
func (a Person) Mine(b Person) int  {
	return mine(a,b)
}


// a相当于this
func mine(a, b Person) int  {
	return a.Age - b.Age
}



// 默认会对每一个参数值进行拷贝
// 可以用其指针而不是对象来声明方法，避免对每个参数值进行拷贝
func (a *Person) Sum2(b Person) int  {
	return a.Age + b.Age
}

