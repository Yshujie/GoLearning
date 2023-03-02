package functionTest

import "fmt"

// 函数声明: 函数名、参数列表、返回值列表、函数体。
// 如果函数没有返回值，则返回列表可以省略。
// 函数从第一条语句开始执行，直到执行return语句或者执行函数的最后一条语句。

// 函数名：TestDeclareFunction
// 参数列表：argus string类型
// 返回值列表：result string 类型
// 函数体：{}

func TestDeclareFunction(argus string) string {
	fmt.Println("There is TestDeclareFunction ")
	fmt.Println("param: ", argus)
	result := "result string"
	fmt.Println("result: ", result)

	return result
}
