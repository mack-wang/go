package main
import (
	"io/ioutil"
	"fmt"
)
func main() {
	// 创建String map
	text, _ := ioutil.ReadFile("./test.txt")
	fmt.Print(string(text))
}