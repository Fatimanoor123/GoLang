package main

import "fmt"

func main() {
    var slc []int
    var number int

    fmt.Println("Enter numbers for the slice (enter -1 when done):")
    
    // Pre-setting a maximum number of inputs for simplicity
    for i := 0; i < 20; i++ { // Assuming the user won't enter more than 20 numbers
        fmt.Scan(&number)
        if number == -1 {
            break // Exit the loop if the sentinel value is encountered
        }
        slc = append(slc, number)
    }

    fmt.Println("Enter the number you want to find:")
    fmt.Scan(&number) // Read the number to find

    found := false // Flag to indicate if the number is found
    index := -1 // Variable to store the index of the found number
    for i := 0; i < len(slc); i++ { // Iterate through the slice
        if slc[i] == number { // Check if current element is the number we're looking for
            found = true
            index = i
            break // Stop the loop if number is found
        }
    }

    if found { // If the number was found, print its index
        fmt.Printf("Number found at index: %d\n", index)
    } else { // If the number was not found, indicate so
        fmt.Println("Number not found in the slice.")
    }
}
