package main

import (
	"fmt"
	"github.com/GoLearning/case/concurrent/goroutineTest"
)

func main() {

	// 并发执行程序
	go goroutineTest.Running()

	// 接收命令行输入
	var input string
	fmt.Scanln(&input)
}
