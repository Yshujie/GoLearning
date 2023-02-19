package main

import (
	"GoLearning/case/hello"
	"GoLearning/case/pointer"
	hungrySingleton "GoLearning/case/singleton/hungry"
)

func main() {

	singletonTest()

}

// 测试单例模式
func singletonTest() {
	var person = hungrySingleton.GetInsOr()

	person.Greet()
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
