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


## 7.创建包，导入包，使用包
1.首先要将你的项目路径添加到$GOPATH中
2.创建包目录
~~~
src/utils/
~~~
3.创建包文件
~~~
src/utils/tempconvs.go
~~~
~~~
// 在tempconvs.go的头部定义包名
package utils
~~~

4.导入包
因为当前项目路径已经添加到$GOPATH,所以可以添加导入包
~~~
import (
	"utils"
)
~~~
5.使用该包下的具体模块
package utils 是在utils目录下的，所以直接使用package名就可以调用该模块,并会在pkg下生成utils.a的编译文件
~~~
	fmt.Println(utils.CToF(80))
~~~


## 8.显示go的相关包的用法
~~~
go doc fmt
~~~

## 9.测试
_test.go为后缀名的源文件并不是go build构建包的一部分，它们是go test测试的一部分。

## 10.函数传参 recieved是值类型还是指针类型 ？
到底是采用值类型还是指针类型主要参考如下原则：

func(w Win) Tally(playerPlayer)int    //w不会有任何改变
func(w *Win) Tally(playerPlayer)int    //w会改变数据

## 11.给import命名一个简短的名字
import 	p2p "gitlab.zhonganonline.com/ann/ann-module/lib/go-p2p"

## 12.无论如何，只有要err都要处理