package main

import "fmt"

func swap(a *int, b *int) {
	*a, *b = *b, *a
}
func main() {
	num1 := 10
	num2 := 20

	fmt.Println("The values before swapping", num1, num2)
	swap(&num1, &num2)

	fmt.Println("The value after swaping", num1, num2)

}
