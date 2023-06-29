package inputValidation

import "strings"

// Validation of Input data.
func PerformValidation(firstName string, lastName string, email string, purchasedTickets int, remaining int) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := purchasedTickets <= int(remaining) && purchasedTickets >= 0
	return isValidName, isValidEmail, isValidTicketNumber
}
