package main

import (
	"fmt"
	"github.com/aadityadike/learn-Go/userInput"
	"github.com/aadityadike/learn-Go/userValidation"
	"strings"
)

// Package level Variables
const TotalTickets uint = 50 // uint simply doesn't take negative elements.
var bookings []string
var remaining = TotalTickets

func main() {
	for {
		// Base case
		if remaining <= 0 {
			fmt.Println("Feel sorry for you! All Tickets are sold for this year, you can comeback next year")
			break
		}

		greetUsers()

		firstName, lastName, email, purchasedTickets := userInput.GetUserInput()

		isValidName, isValidEmail, isValidTicketNumber := userValidation.PerformValidation(firstName, lastName, email, purchasedTickets, int(remaining))

		// Merging first and last name.
		bookings = append(bookings, firstName+" "+lastName)

		if isValidTicketNumber && isValidEmail && isValidName {
			takeYourTicket(email, purchasedTickets)
		} else {
			validationErrors(isValidEmail, isValidName, isValidTicketNumber)
			continue
		}
	}
}

// greeting users.
func greetUsers() {
	fmt.Println("\nWelcome to Kubecon(2023)")
	fmt.Printf("Get your Tickets here! we have %v Tickets \n", remaining)
}

// It will bring First Name for you.
func getFirstName() string {
	// Dividing first and last name.
	firstNames := []string{}
	// there are two fields required here, index & the array or slice but since we are not using index we can put underscore here.
	for _, bookings := range bookings {
		/* strings.Fields is function provided by string, this is use to separate two string with spaces it will look something like array or slice [firstPartOfString, secondPartOfString]
		 */
		var names = strings.Fields(bookings)
		var updatedFirstName = names[0]
		firstNames = append(firstNames, updatedFirstName)
	}
	return firstNames[len(firstNames)-1]
}

// It will bring you a Ticket & also reduce the number of remaining tickets from available tickets.
func takeYourTicket(email string, purchasedTickets int) {
	// It will bring First Name for you.
	userName := getFirstName()

	fmt.Printf("Yeah, %v you are coming to this year's event Happy to have you!\n", userName)
	fmt.Printf("Your %v Tickets will be sent on %v\n", purchasedTickets, email)

	remaining -= uint(purchasedTickets)
}

// It will give errors, if there is error by user.
func validationErrors(isValidEmail bool, isValidName bool, isValidTicketNumber bool) {
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
