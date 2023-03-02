package simple

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Println("Hello, I'm ", p.Name)
	fmt.Println("There is simple factory")
}
