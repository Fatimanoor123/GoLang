package main

import "fmt"

func main() {
	//var n int
	var value int
	var max int
	var hasInput bool = false

	fmt.Println("Enter numbers (type 0 to finish):")

	for {
		fmt.Scan(&value)
		if value == 0 {
			break
		}
		if !hasInput || value > max {
			max = value
			hasInput = true
		}
	}

	if !hasInput {
		fmt.Println("No numbers entered.")
	} else {
		fmt.Printf("The maximum value is: %d\n", max)
	}
}
