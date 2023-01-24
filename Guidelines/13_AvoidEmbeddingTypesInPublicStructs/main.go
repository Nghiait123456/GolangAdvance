package main

import "fmt"

type AbstractList struct{}
type Entity struct {
}

// Add adds an entity to the list.
func (l *AbstractList) Add(e Entity) {
	fmt.Println("add")
}

// Remove removes an entity from the list.
func (l *AbstractList) Remove(e Entity) {
	fmt.Println("Remove")
}

// ConcreteListGood is a list of entities.
type ConcreteListBad struct {
	*AbstractList
}

// ConcreteListGood is a list of entities.
type ConcreteListGood struct {
	list *AbstractList
}

// Add adds an entity to the list.
func (l *ConcreteListGood) Add(e Entity) {
	l.list.Add(e)
}

// Remove removes an entity from the list.
func (l *ConcreteListGood) Remove(e Entity) {
	l.list.Remove(e)
}

func main() {
	fmt.Println("Embedded type is rarely necessary. Although writing these delegate methods is tedious.The additional effort hides an implementation detail, leaves more opportunities for change, and also eliminates indirection for discovering the full List interface in documentation.")
	fmt.Println("This is bad embedded AbstractList")
	cBad := ConcreteListBad{}
	cBad.Add(Entity{})
	cBad.Remove(Entity{})

	fmt.Println("This is good embedded AbstractList")
	cGood := ConcreteListGood{
		list: &AbstractList{},
	}
	cGood.Add(Entity{})
	cGood.Remove(Entity{})
}
