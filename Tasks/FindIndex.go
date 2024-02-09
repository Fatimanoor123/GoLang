//find maximum value of integer in slice

package main

import "fmt"

func main() {
	var slc []int
	var element int
	slc = append(slc, 21)
	slc = append(slc, 12)
	slc = append(slc, 10)
	slc = append(slc, 1)
	slc = append(slc, 14)
	fmt.Println("The elements in Slice are: ", slc)
	fmt.Println("Do you want to check the index of any element, then enter element")

	fmt.Scanln(&element)

	for i := 0; i < 5; i++ {
		if slc[i] == element {
			fmt.Println("The address of desired element is as follows: ", &slc[i])1
		}
	}

}
