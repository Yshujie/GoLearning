package arrayType

import "fmt"

var PublicArr = [...]int{1, 2, 3}
var PublicArr2 = [...]int{}

// package 内可见
var arr0 [5]int
var arr1 [5]int = [5]int{1, 2, 3}
var arr2 = [5]int{1, 2, 3, 4, 5}
var arr3 = [...]int{1, 2, 3, 4, 5, 6}
var strArr = [5]string{3: "hello", 4: "world"}

func Test1() {
	fmt.Println("arr0: ", arr0)
	arr0[0] = 3
	arr0[1] = 2
	arr0[2] = 1
	fmt.Println("arr0: ", arr0)

	fmt.Println("arr1: ", arr1)
	fmt.Println("arr2: ", arr2)
	fmt.Println("arr3: ", arr3)
	fmt.Println("strArr: ", strArr)

	PublicArr[0] = 3
}
