package lazy_unsafety

import "fmt"

type Person struct {
	name string
	sex  string
	age  int
}

func (person Person) Greet() {
	fmt.Println("Hi~ There is unSafety lazy model. I'm ", person.name, ", and I'm a ", person.sex, ". In this years, I'm ", person.age, "years old")
}
