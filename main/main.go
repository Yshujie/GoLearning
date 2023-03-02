package main

import "github.com/GoLearning/case/factory/simple"

func main() {

	testSimpleFactory()

}

func testSimpleFactory() {
	person := simple.NewPerson("moons", 3)

	person.Greet()
}
