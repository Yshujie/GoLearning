package arrayType

import "fmt"

var PublicArr = [...]int{1, 2, 3}
var PublicArr2 = [...]int{}

// package 内可见
var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var strArr = [5]string{3: "hello", 4: "world"}

func Test1() {
	fmt.Println(arr0)
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(strArr)

	PublicArr[0] = 3
}
