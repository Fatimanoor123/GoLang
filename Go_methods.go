package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height

}
func main() {
	s := rect{height: 2, width: 4}
	fmt.Println("the value of area of rectangle is ", s.area())

}
