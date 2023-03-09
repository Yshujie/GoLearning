package main

import (
	"fmt"
	"github.com/GoLearning/case/initTest"
)

func main() {

	initTest.SayHello()

}

func init() {
	fmt.Println("in main init 1")
}

func init() {
	fmt.Println("in main init 2")
}
