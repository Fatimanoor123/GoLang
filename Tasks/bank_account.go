package main

import (
	"fmt"
)

// Define the BankAccount struct
type BankAccount struct {
	accountNumber string
	balance       float64
}

// Method to deposit money into the account
func (ba *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		ba.balance += amount
		fmt.Printf("Deposited $%.2f, new balance is $%.2f\n", amount, ba.balance)
	} else {
		fmt.Println("Please enter a positive amount to deposit.")
	}
}

// Method to withdraw money from the account
func (ba *BankAccount) Withdraw(amount float64) {
	if amount > 0 && amount <= ba.balance {
		ba.balance -= amount
		fmt.Printf("Withdrew $%.2f, new balance is $%.2f\n", amount, ba.balance)
	} else if amount > ba.balance {
		fmt.Println("Insufficient funds for this withdrawal.")
	} else {
		fmt.Println("Please enter a positive amount to withdraw.")
	}
}

func main() {
	// Create a new BankAccount instance
	account := BankAccount{accountNumber: "123456789", balance: 100.00}

	// Display initial account balance
	fmt.Printf("Initial balance is $%.2f\n", account.balance)

	// Deposit money
	account.Deposit(50.00)

	// Withdraw money
	account.Withdraw(25.00)

	// Attempt to withdraw more than the balance
	account.Withdraw(200.00)
}
