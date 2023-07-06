package inputValidation

import (
	"fmt"
	"strings"
)

// Validation of Input data.
func PerformValidation(firstName string, lastName string, email string, purchasedTickets int, remaining int) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := purchasedTickets <= int(remaining) && purchasedTickets >= 0
	return isValidName, isValidEmail, isValidTicketNumber
}

// It will give errors, if there is error by user.
func ValidationErrors(isValidEmail bool, isValidName bool, isValidTicketNumber bool) {
	if !isValidEmail {
		fmt.Println("Email address that you enter doesn't contain @ sign")
	}
	if !isValidName {
		fmt.Println("Please enter First & Last Name properly")
	}
	if !isValidTicketNumber {
		fmt.Println("You enter invalid Ticket number")
	}
}