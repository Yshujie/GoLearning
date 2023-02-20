package main

import (
	"fmt"
	lazy_safety "github.com/GoLearning/case/singleton/lazy.safety"
	lazy_unsafety "github.com/GoLearning/case/singleton/lazy.unsafety"
)

func main() {

	lazyUnSafetyTest()

	lazySafetyTest()

}

func lazyUnSafetyTest() {
	var person1 = lazy_unsafety.GetInsOr()
	person1.Greet()

	var person2 = lazy_unsafety.GetInsOr()
	person2.Greet()

	fmt.Println(&person1)
	fmt.Println(&person2)
}

func lazySafetyTest() {
	var person1 = lazy_safety.GetInsOr()
	person1.Greet()

	var person2 = lazy_safety.GetInsOr()
	person2.Greet()

	fmt.Println(&person1)
	fmt.Println(&person2)
}
