package userInput

import "fmt"

// getting user input.
func GetUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var purchasedTickets int
	var email string

	fmt.Println("Enter First name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter Last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter email: ")
	fmt.Scan(&email)
	fmt.Println("Enter the number of ticket you want to purchase: ")
	fmt.Scan(&purchasedTickets)

	return firstName, lastName, email, purchasedTickets
}
