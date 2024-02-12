//write a recursive function in go to calculate factorial of non-negative number

package main

import "fmt"

// Recursive function to calculate factorial
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    var n int
    fmt.Print("Enter a non-negative integer: ")
    fmt.Scan(&n)
    fmt.Printf("Factorial of %d is %d\n", n, factorial(n))
}
