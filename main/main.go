package main

import (
	"GoLearning/case/hello"
	"GoLearning/case/pointer"
	hungrySingleton "GoLearning/case/singleton/hungry"
	"fmt"
)

func main() {

	printTest()

	helloWorldTest()

}

func printTest() {
	fmt.Println("test print something. main package needs to import fmt package, event if the fmt package has already imported by hello package")
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
