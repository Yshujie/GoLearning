package main

import (
	"fmt"
)

func main() {

	printTest()

}

func printTest() {
	fmt.Println("test print something. main package needs to import fmt package, event if the fmt package has already imported by hello package")
}
