package hellomock

import "fmt"

// Person Person
type Person struct {
	name string
}

// NewPerson NewPerson
func NewPerson(name string) *Person {
	return &Person{
		name: name,
	}
}

// SayHello SayHello
func (p *Person) SayHello(name string) (word string) {
	return fmt.Sprintf("Hello %s, welcome come to our company, my name is %s", name, p.name)
}
