package main

import "fmt"

// Book struct with a Publisher field
type book struct {
	Title     string
	Isbn      string
	Price     float64
	Publisher string
	Quantity  int
}

// Initialize Books slice with some books
var Books = []book{
	{"The Go Programming Language", "978-0134190440", 46.75, "Addison-Wesley", 10},
	{"Go in Action", "978-1617291784", 29.99, "Manning Publications", 5},
	{"Learning Go", "978-1492077213", 35.99, "O'Reilly Media", 8},
}

type user struct {
	name, password, email string
}

var users []user

// Function to display all books
func displayBooks() {
	if len(Books) == 0 {
		fmt.Println("No books available.")
		return
	}
	fmt.Println("Available books:")
	for _, book := range Books {
		fmt.Printf("Title: %s, ISBN: %s, Price: $%.2f, Publisher: %s, Quantity: %d\n", book.Title, book.Isbn, book.Price, book.Publisher, book.Quantity)
	}
}

// Function to add books and display the updated list
func addBooks(title, isbn string, price float64, publisher string, quantity int) {
	Books = append(Books, book{Title: title, Isbn: isbn, Price: price, Publisher: publisher, Quantity: quantity})
	fmt.Println("Book added successfully. Current list of books:")
	displayBooks()
}

// Function for deleting books by ISBN and displaying the updated list
func delBooks(isbn string) {
	found := false
	for i, book := range Books {
		if book.Isbn == isbn {
			Books = append(Books[:i], Books[i+1:]...)
			found = true
			break
		}
	}

	if found {
		fmt.Println("Book deleted successfully. Updated list of books:")
		displayBooks()
	} else {
		fmt.Println("Book with ISBN", isbn, "not found.")
	}
}

// Function to update books by ISBN
func updBooks(isbn string) {
	found := false
	for i, _ := range Books {
		if Books[i].Isbn == isbn {
			fmt.Print("Enter new title: ")
			fmt.Scan(&Books[i].Title)
			fmt.Print("Enter new price: ")
			fmt.Scan(&Books[i].Price)
			fmt.Print("Enter new publisher: ")
			fmt.Scan(&Books[i].Publisher)
			fmt.Print("Enter new quantity: ")
			fmt.Scan(&Books[i].Quantity)
			found = true
			fmt.Println("Book updated successfully. Updated list of books:")
			displayBooks()
			break
		}
	}
	if !found {
		fmt.Println("Book with ISBN", isbn, "not found.")
	}
}
func searchBook(isbn string) {
	found := false
	for _, book := range Books {
		if book.Isbn == isbn {
			fmt.Printf("Book found: Title: %s, ISBN: %s, Price: $%.2f, Publisher: %s, Quantity: %d\n", book.Title, book.Isbn, book.Price, book.Publisher, book.Quantity)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Book not Found")
	}
}

func main() {
	var name, password, email string
	fmt.Println("Welcome to Cagreg Book Inventory!")
	fmt.Println("Before getting started, please enter your name, password and email address")
	fmt.Scan(&name, &password, &email)

	for {
		var en int
		fmt.Println("\n1_ Display Books \t", "2_ Add Books \t \n", "3_ Delete Books \t", "4_ Update Books \t \n", "5_ Search \t ", "6_Exit")
		fmt.Scanln(&en)
		switch en {
		case 1:
			displayBooks()
		case 2:
			var title, isbn, publisher string
			var price float64
			var quantity int
			fmt.Print("Enter book title: ")
			fmt.Scan(&title)
			fmt.Print("Enter ISBN: ")
			fmt.Scan(&isbn)
			fmt.Print("Enter price: ")
			fmt.Scan(&price)
			fmt.Print("Enter publisher: ")
			fmt.Scan(&publisher)
			fmt.Print("Enter quantity: ")
			fmt.Scan(&quantity)
			addBooks(title, isbn, price, publisher, quantity)
		case 3:
			var isbn string
			fmt.Print("Enter the ISBN of the book to delete: ")
			fmt.Scan(&isbn)
			delBooks(isbn)
		case 4:
			var isbn string
			fmt.Print("Enter the ISBN of the book to update: ")
			fmt.Scan(&isbn)
			updBooks(isbn)
		case 5:
			var isbn string
			fmt.Print("Enter the title to search a book ")
			fmt.Scan(isbn)
			searchBook(isbn)
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
