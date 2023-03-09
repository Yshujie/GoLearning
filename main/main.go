package main

import (
	"fmt"
	"github.com/GoLearning/case/dataType/arrayType"
)

func main() {

	testArrayDataType()

}

func testArrayDataType() {

	// arrayType 包向外暴露的变量
	fmt.Println(arrayType.PublicArr)
	fmt.Println(len(arrayType.PublicArr))
	fmt.Println(arrayType.PublicArr2)
	fmt.Println(len(arrayType.PublicArr2))

	// arrayType 包私有的变量
	// arrayType.arr0 访问不到

	// arrayType 包的函数
	arrayType.Test1()

	fmt.Println(arrayType.PublicArr)

}
