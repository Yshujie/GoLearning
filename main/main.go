package main

import (
	"fmt"
	"github.com/GoLearning/case/hello"
)

func main() {

	printTest()

	helloTest()

}

func printTest() {
	fmt.Println("test print something. main package needs to import fmt package, event if the fmt package has already imported by hello package")
}

func helloTest() {
	hello.Hello()
}
