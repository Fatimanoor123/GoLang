package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Println("Do you want to know the month then enter any number between 1-12")
	fmt.Scanln(&num)

	switch num {

	case 1:
		fmt.Println("It's January")
	case 2:
		fmt.Println("It's February")
	case 3:
		fmt.Println("It's march")
	case 4:
		fmt.Println("It's April")
	case 5:
		fmt.Println("It's May")
	case 6:
		fmt.Println("It's June")
	case 7:
		fmt.Println("It's July")
	case 8:
		fmt.Println("It's August")
	case 9:
		fmt.Println("It's September")
	case 10:
		fmt.Println("It's October")
	case 11:
		fmt.Println("It's November")
	case 12:
		fmt.Println("It's December")

	}

}
