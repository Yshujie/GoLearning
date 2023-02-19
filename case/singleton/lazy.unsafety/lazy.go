package lazy_unsafety

// 懒汉模式 -- 并发不安全

var person *Person

func GetInsOr() *Person {
	if person == nil {
		person = &Person{
			"Jarvis",
			"boy",
			3,
		}
	}

	return person
}
