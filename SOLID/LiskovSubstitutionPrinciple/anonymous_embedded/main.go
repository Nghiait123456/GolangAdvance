package main

import "fmt"

type PersonInterface interface {
	GetName() bool
}

type Person struct {
	name string
}

func (p *Person) GetName() bool {
	fmt.Println(p.name)
	return true
}

type Student struct {
	Person
}

type Teacher struct {
	Person
}

type Identification struct {
	p PersonInterface
}

func main() {
	t := Teacher{
		Person{
			name: "Teacher 1",
		},
	}
	i1 := Identification{
		p: &t,
	}
	i1.p.GetName()

	s := Student{
		Person{
			name: "Student 1",
		},
	}
	i2 := Identification{
		p: &s,
	}
	i2.p.GetName()
}
