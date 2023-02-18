package pointerTest

import "fmt"

// 测试1，对指针类型进行测试

func Test1()  {
	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b

	fmt.Println("a = ", a)  // 1
	fmt.Println("&a = ", &a) // 变量 a 的地址值
	fmt.Println("*&a = ", *&a) // 1
	fmt.Println("b = ", b) // 变量 a 的地址值
	fmt.Println("&b = ", &b) // 变量 b 的地址值
	fmt.Println("*&b = ", *&b) // 变量 a 的地址值
	fmt.Println("*b = ", *b) // 1
	fmt.Println("c = ", c) // 变量 b 的地址值
	fmt.Println("*c = ", *c) // 变量 a 的地址值
	fmt.Println("&c = ", &c) // 变量 c 的地址值
	fmt.Println("*&c = ", *&c) // 变量 b 的地址值
	fmt.Println("**c = ", **c) // 1
	fmt.Println("***&*&*&*&c = ", ***&*&*&*&*&c) // 1
	fmt.Println("x = ", x) // 1
}

// 测试2，对指针的指针类型进行测试

func Test2()  {
	var a int = 1
	var b *int = &a
	// var x *string = &a 错误，cannot use &a (value of type *int) as type *string in variable declaration
	var c string = "1"
	var d *string = &c

	fmt.Println("a: ", a )
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)
}