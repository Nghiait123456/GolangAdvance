package main

import "fmt"

type PersonInterface interface {
	PrintName()
}

type Person struct {
	name string
}

func (p *Person) PrintName() {
	fmt.Println(p.name)
}

type Student struct {
	p     PersonInterface
	class string
	point string
}

type StudentInterface interface {
	PersonInterface
	GetStudentClass() string
	GetStudentPoint() string
}

func (t *Student) PrintName() {
	t.p.PrintName()
}

func (t *Student) GetStudentClass() string {
	return t.class
}

func (t *Student) GetStudentPoint() string {
	return t.point
}

type Teacher struct {
	p     PersonInterface
	class string
	point string
}

func (t *Teacher) PrintName() {
	t.p.PrintName()
}

func (t *Teacher) GetTeacherClass() string {
	return t.class
}

func (t *Teacher) GetTeacherPoint() string {
	return t.point
}

type TeacherInterface interface {
	PersonInterface
	GetTeacherClass() string
	GetTeacherPoint() string
}

type Identification struct {
	p PersonInterface
}

type StudentIdentification struct {
	p StudentInterface
}

type TeacherIdentification struct {
	p TeacherInterface
}

func main() {
	p1 := Person{
		name: "AAA",
	}
	p2 := Person{
		name: "BBB",
	}

	t := Teacher{
		p:     &p1,
		point: "A",
		class: "B",
	}

	s := Student{
		p:     &p2,
		point: "A",
		class: "B",
	}

	i := Identification{
		p: &t,
	}

	i.p.PrintName()

	//student := StudentIdentification{
	//	p: &t,
	//}

	//teacher := StudentIdentification{
	//	p: &s,
	//}

	student := StudentIdentification{
		p: &s,
	}
	student.p.PrintName()
}
