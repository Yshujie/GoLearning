package hungrySingleton

import "fmt"

type Person struct {
	name string
	sex  string
	age  int
}

func (p Person) Greet() {
	fmt.Println("Hi~ There is hungry singleton. My name is ", p.name, ", I'm", p.age, " years old, and I'm a ", p.sex)
}
