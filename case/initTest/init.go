package initTest

import "fmt"

func init() {
	fmt.Println("in initTest init 1")
}

func init() {
	fmt.Println("in initTest init 2")
}

func SayHello() {
	fmt.Println("Hello ~~~ ")
}
