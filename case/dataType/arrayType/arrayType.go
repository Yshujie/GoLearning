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

// Test1 测试 package 内声明的数组变量
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

// Test2 测试函数内部声明的数组变量
func Test2() {
	arr0 := [3]int{3, 2, 1}
	a := [3]int{1, 2}
	b := [...]int{1, 2, 3, 4}
	c := [5]int{1: 9527, 3: 579}
	d := [...]struct {
		name string
		age  int
	}{
		{"Jarvis", 5},
		{"moons", 2},
	}

	fmt.Println("arr0: ", arr0)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)
}

// Test3 测试一维数组遍历
func Test3() {
	arr := [...]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	fmt.Println(" ---- start: for 循环遍历 ---- ")
	iterateArrayOfFor(arr)
	fmt.Println(" ----  end : for 循环遍历 ---- ")

	fmt.Println(" ---- start: range 循环遍历 ---- ")
	iterateArrayOfRange(arr)
	fmt.Println(" ----  end : range 循环遍历 ---- ")
}

// 测试 for 循环遍历数组
func iterateArrayOfFor(arr [10]int) {
	for i := 0; i < len(arr); i++ {
		println("index:", i, "; value:", arr[i])
	}
}

// 测试 rang 循环遍历数组
func iterateArrayOfRange(arr [10]int) {
	for index, value := range arr {
		fmt.Println("index:", index, "; value:", value)
	}
}
