package lazy_safety

import "fmt"

type Person struct {
	name string
	sex  string
	age  int
}

func (person Person) Greet() {
	fmt.Println("Hi~ There is lazy safety singleton. I'm ", person.name, "and I'm a ", person.sex, ", and in this years, I'm ", person.age, "years old")
}
