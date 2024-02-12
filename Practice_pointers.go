package main

import "fmt"

func zeroval(ival int) {
	ival = 0
	fmt.Println("Inside zero value function")
}

func zeroptr(iptr *int) {
	*iptr = 0

}
func main() {
	i := 0
	fmt.Println("The value of i is:", i)
	zeroval(i)
	fmt.Println("after calling zero value")
	zeroptr(&i)
	fmt.Println("the latest value of i is:", i)

}
