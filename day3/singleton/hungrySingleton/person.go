package hungrySingleton

import "fmt"

type person struct {
	Name string
	Sex string
	Age int
}

func (p person) Greet() {
	fmt.Println("Hi~ There is hungry singleton. My name is %s, I'm %s age years old", p.Name, p.Age)
}