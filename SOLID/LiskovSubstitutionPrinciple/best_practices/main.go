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
	name  string
	class string
	point string
}

type StudentInterface interface {
	PersonInterface
	GetStudentClass() string
	GetStudentPoint() string
}

func (t *Student) PrintName() {
	fmt.Println("name = ", t.name)
}

func (t *Student) GetStudentClass() string {
	return t.class
}

func (t *Student) GetStudentPoint() string {
	return t.point
}

type Teacher struct {
	name  string
	class string
	point string
}

func (t *Teacher) PrintName() {
	fmt.Println("name = ", t.name)
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
	t := Teacher{
		name:  "teacher",
		point: "A",
		class: "B",
	}

	s := Student{
		name:  "student",
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
