package main

import (
	"fmt"
	"github.com/GoLearning/case/processControl/functionTest"
)

func main() {

	testFunction()

}

func testFunction() {

	// 直接传入匿名函数
	result := functionTest.TestFuncParams(func() int {
		return 9527
	})

	fmt.Println("result: ", result)

	result2 := functionTest.TestFormat(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)

	fmt.Println("result2: ", result2)

}
