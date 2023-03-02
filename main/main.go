package main

import "github.com/GoLearning/case/factory/abstract"

func main() {

	testAbstractFactory()

}

func testAbstractFactory() {
	ai := abstract.NewPerson("Jarvis", 8)

	ai.Greet("I'm your assistant")
}
