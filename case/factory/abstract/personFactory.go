package abstract

// 工厂模型 -- 抽象工厂

// 抽象工厂和简单工厂唯一的区别，是它返回的是接口而不是结构体

func NewPerson(name string, age int) Person {
	return person{
		name: name,
		age:  age,
	}
}
