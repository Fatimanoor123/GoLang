package main

import "fmt"

func add(a int, b int) int {
	fmt.Println("The values for addition are ", a, b)
	return a + b
}

func sub(a int, b int) int {
   fmt.Println("Values for subtraction is: ", a , b)
	return a - b

}
func main() {

	result := add(3, 5)
	fmt.Println("The addition of two numbers is: ", result)

	res := sub(5, 1)

	fmt.Println("The sub of two number is: ", res)

}
