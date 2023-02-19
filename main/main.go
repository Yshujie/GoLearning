package main

import (
	"GoLearning/case/hello"
	"GoLearning/case/pointer"
)

func main() {

	pointerTest()

}

// 跑通 hello world
func helloWorldTest() {
	hello.Hello()
}

// 测试 Go 语言指针类型
func pointerTest() {
	pointer.Test1()

	pointer.Test2()
}
