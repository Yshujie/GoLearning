package main

import (
	"fmt"
	"github.com/GoLearning/case/processControl/functionTest"
)

func main() {

	testFunction()

}

func testFunction() {
	functionTest.TestDeclareFunction1("Hello")

	result1, result2 := functionTest.TestDeclareFunction2("Hello", 9527)
	fmt.Println("I got result1: ", result1)
	fmt.Println("I got result2: ", result2)
}
