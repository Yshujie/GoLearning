package lazy_safety

import "sync"

var person *Person
var mu sync.Mutex

func GetInsOr() *Person {
	if nil == person {
		mu.Lock()

		if nil == person {
			person = &Person{
				"Jarvis",
				"boy",
				3,
			}
		}

		mu.Unlock()
	}

	return person
}
