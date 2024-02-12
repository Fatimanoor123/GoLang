package main

import "fmt"

// Define the Person struct
type Person struct {
	name  string
	age   int
	email string
}

// Function to create a new Person instance
func NewPerson(name string, age int, email string) Person {
	return Person{name: name, age: age, email: email}
}

func main() {
	// Creating a new Person instance
	person1 := NewPerson("John Doe", 30, "johndoe@example.com")

	// Printing the Person details
	fmt.Println("Person Details:")
	fmt.Printf("Name: %s\n", person1.name)
	fmt.Printf("Age: %d\n", person1.age)
	fmt.Printf("Email: %s\n", person1.email)
}
