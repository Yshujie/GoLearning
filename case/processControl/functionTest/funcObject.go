package functionTest

// 函数是第一类对象，可作为参数传递。 可以将负责签名定义为函数类型，以便于阅读

// TestFuncParams 测试 -- 函数作为参数传入
func TestFuncParams(fn func() int) int {
	return fn()
}

// FormatFunc 定义函数类型
// 这里其实是在定义函数的声明（参数列表、返回值列表），已明确按指定的规范传进来函数
type formatFunc func(s string, x, y int) string

// TestFormat 测试 -- 定义 Format 函数
func TestFormat(fn formatFunc, s string, x, y int) string {
	return fn(s, x, y)
}
