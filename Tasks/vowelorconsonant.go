package main

import "fmt"

func main() {
	var ch byte
	fmt.Println("Enter a character to see if it is vowel or not")
	fmt.Scanln(&ch)

	switch ch {
	case 'a':
		fmt.Println("It's a vowel")
	case 'e':
		fmt.Println("It's a vowel")
	case 'i':
		fmt.Println("It's a vowel")
	case 'o':
		fmt.Println("It's a vowel")
	case 'u':
		fmt.Println("It's a vowel")
	default:
		fmt.Println("You enter consonant not a vowel")
	}

}
