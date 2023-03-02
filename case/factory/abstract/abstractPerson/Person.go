package abstractPerson

import "fmt"

type Person interface {
	Greet(welcomeMsg string)
}

type person struct {
	name string
	age  int
}

func (p person) Greet(welcomeMsg string) {
	fmt.Println("Hello, I'm ", p.name)
	fmt.Println("There is abstract factory")
	fmt.Println(welcomeMsg)
}
