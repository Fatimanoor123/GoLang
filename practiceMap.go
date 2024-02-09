package main

import "fmt"

func main() {
	nums := []int{1, 2, 122, 423, 12}
	sum:= 0
	for _, num:= range nums {
		sum += num

	}
	fmt.Println("The sum of each element in slice is as follows: ", sum)
    
	mp:= map[string]string{"a1" : "apple", "b ": "banana", "c" :"cat"}

	for k , v:= range mp{
		fmt.Println("%s--> %s \n", k, v)
	}
		
	}
