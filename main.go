package main

import (
	"fmt"
	"sync"
	"time"
	
	"github.com/aadityadike/BookYourTicketNow/OtherConcepts"
	"github.com/aadityadike/BookYourTicketNow/inputValidation"
	"github.com/aadityadike/BookYourTicketNow/userInput"
)

// Package level Variables
const TotalTickets uint = 50 // uint simply doesn't take negative elements.
var remaining = TotalTickets
var waitGroup = sync.WaitGroup{}

// var bookings []string
// slice of Maps.
// var bookings = make([]map[string]string, 0)

// Slice of Struct - Structure
var bookings = make([]UserData, 0)

// type basically means we are creating a new type of data. Now UserData is new data type.
type UserData struct {
	FirstName        string
	LastName         string
	email            string
	purchasedTickets int
}

func main() {
	// for {
	// Base case

	

	if remaining <= 0 {
		fmt.Println("Feel sorry for you! All Tickets are sold for this year, you can comeback next year")
		// break
	}

	greetUsers()

	firstName, lastName, email, purchasedTickets := userInput.GetUserInput()

	isValidName, isValidEmail, isValidTicketNumber := inputValidation.PerformValidation(firstName, lastName, email, purchasedTickets, int(remaining))

	// Create a Map for User details.
	/*
		var userData = make(map[string]string)
		userData["FirstName"] = firstName
		userData["LastName"] = lastName
		userData["Email"] = email
		userData["PurchasedTickets"] = strconv.FormatInt(int64(purchasedTickets), 10)
	*/

	var userData = UserData{
		FirstName:        firstName,
		LastName:         lastName,
		email:            email,
		purchasedTickets: purchasedTickets,
	}

	// Merging first and last name.
	bookings = append(bookings, userData)
	fmt.Println(bookings)

	if isValidTicketNumber && isValidEmail && isValidName {
		takeYourTicket(purchasedTickets)
		/*
		* go keyword starts a new goroutine.
		* A goroutine is a lightweight thread managed by the Go runtime.
		 */
		waitGroup.Add(1)
		go sendTicketToUser(email, purchasedTickets)
	} else {
		inputValidation.ValidationErrors(isValidEmail, isValidName, isValidTicketNumber)
		// continue
	}
	waitGroup.Wait()
}

// }

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
		firstNames = append(firstNames, bookings.FirstName)
	}
	return firstNames[len(firstNames)-1]
}

// It will bring you a Ticket & also reduce the number of remaining tickets from available tickets.
func takeYourTicket(purchasedTickets int) {
	// It will bring First Name for you.
	userName := getFirstName()

	fmt.Printf("Yeah, %v you are coming to this year's event Happy to have you!\n", userName)
	remaining -= uint(purchasedTickets)
}

func sendTicketToUser(email string, purchasedTickets int) {
	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf("Your %v Tickets will be sent on %v\n", purchasedTickets, email)
	fmt.Println("##############################")
	fmt.Println(ticket) // you can return value ticket also.
	fmt.Printf("Ticket is sent to %v\n", email)
	fmt.Println("##############################")
	waitGroup.Done()
}
