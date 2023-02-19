package main

import lazy_unsafety "github.com/GoLearning/case/singleton/lazy.unsafety"

func main() {

	lazyUnSafetyTest()

}

func lazyUnSafetyTest() {
	var person = lazy_unsafety.GetInsOr()

	person.Greet()
}
