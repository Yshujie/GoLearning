package hungry

// 单例模式 -- 饿汉方式
// 注意：因为实例是在包被导入时初始化的，所以如果初始化耗时，会导致程序加载时间比较长。

var p *Person = &Person{
	"Jevers",
	"boy",
	3,
}

func GetInsOr() *Person {
	return p
}
