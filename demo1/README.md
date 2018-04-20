## 1.执行文件
~~~
go run xxx.go
~~~

## 2.编译文件
~~~
go build xxx.go
~~~

## 3.从远程获取程序并安装到gopath目录下
~~~
go get gopl.io/ch1/helloworld
~~~

## 4.go不写分号，换行符会自动变成分号

## 5.可以使用_下划线来用做占位符，占据变量位置

~~~
// 下面要求返回 content, error，但error我不想要，所以用_占位
text, _ := ioutil.ReadFile("./test.txt")
~~~

## 6.变量类型
var、const、type和func，分别对应变量、常量、类型和函数实体对象的声明