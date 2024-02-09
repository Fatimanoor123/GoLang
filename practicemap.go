package main

import (
	"fmt"
)

func main() {
	// Map functions

	//declaration of maps
	m := make(map[string]int)

	//initialization of map
	m["Fatima"] = 25
	m["Ayesha"] = 12
	m["Noor"] = 10

	fmt.Println("The values in map is: ", m)

	//Initialization and declarltion

	f := map[string]int{"k1": 7, "K2": 13, "k3": 15}
	fmt.Println("The second map function is : ", f)
	// deleting the values from map functions

	delete(m, "Ayesha")

	fmt.Println("The value of updated m map function is as follows: ", m)

	delete(f, "k3")
	fmt.Println("The updated value for map function is: ", f)
}
