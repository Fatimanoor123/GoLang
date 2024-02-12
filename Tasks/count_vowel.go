package main

import "fmt"

// Function to count vowels in a string
func countVowels(s string) int {
    count := 0
    // Define the vowels
    vowels := "aeiouAEIOU"
    // Loop through each character in the string
    for i := 0; i < len(s); i++ {
        // Check if the character is in the vowels string
        for j := 0; j < len(vowels); j++ {
            if s[i] == vowels[j] {
                count++
                break
            }
        }
    }
    return count
}

func main() {
    // Prompt the user for a string
    fmt.Println("Enter a string:")
    var input string
    fmt.Scanln(&input)
    // Count the vowels in the string
    fmt.Println("Number of vowels:", countVowels(input))
}
