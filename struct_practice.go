package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 55
	return &p
}

func main() {
	fmt.Println(person{name: "Misbah", age: 24})
	fmt.Println(person{name: "Fatima", age: 25})
	fmt.Println(newPerson("Uncle"))

}
