package main

import (
	"os"
	"text/template"
)

type Man2 struct {
	Name string
	Age  int
}

func main() {
	// 定义变量
	bob := Man2{"bob", 23}

	// 定义模版，和变量替换的位置
	tmpl, _ := template.New("test").Parse("hello, {{.Name}} {{.Age}}") //建立一个模板，内容是"hello, {{.}}"

	// 输出
	_ = tmpl.Execute(os.Stdout, bob) //将string与模板合成，变量name的内容会替换掉{{.}}

}
