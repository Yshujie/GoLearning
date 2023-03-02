package simple

// 工厂模式 -- 简单工厂

// 简单工厂的优点：简单工厂模式可以确保我们创建的实例具有需要的参数，进而保证实例的方法可以按预期执行。

func NewPerson(name string, age int) *Person {
	return &Person{
		name,
		age,
	}
}
