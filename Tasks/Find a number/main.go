// find maximum value in slice
package main

import (
	"fmt"
	"sort"
)

func main() {
	var a []int
	a = append(a, 23)
	a = append(a, 2)
	a = append(a, 311)
	a = append(a, 42)
	a = append(a, 51)
	fmt.Println("Values in slice are", a)
	sort.Ints(a)
	fmt.Println("the maximum value is: ", a[len(a)-1])

}
