//Question find sum of all integer of an array


package main

import "fmt"

func main() {

	var j int
	j = 0
	arr := [5]int{23, 3, 12, 11, 45}
	for i := 0; i < 5; i++ {
		j = j + arr[i]
	}
	fmt.Println("The sum of entire array is as follow: ", j)

}

