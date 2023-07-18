package userInput

import "fmt"

// Example of Switch statement.
func Switch() {
	// let's take a example, we want to determine in which city user wants to attend event
	// we are assuming that event is virtual and held at various cities
	var city string
	// user input
	fmt.Println("Enter a city where you want to attend a event")
	fmt.Println("The Event is hosted at following cities:\nNew York\nSingapore\nIndia\nUnited Kingdom\nAmsterdam\nManchester\nMumbai")
	fmt.Println("Choose one of them: ")
	fmt.Scan(city)

	switch city {
	case "New York":
		fmt.Printf("Congratulations! you are going to attend Kubecon from %v ", city)
	case "Singapore":
		fmt.Printf("Congratulations! you are going to attend Kubecon from %v ", city)
	case "India", "Mumbai":
		fmt.Printf("Congratulations! you are going to attend Kubecon from %v ", city)
	case "United Kingdom":
		fmt.Printf("Congratulations! you are going to attend Kubecon from %v ", city)
	case "Amsterdam":
		fmt.Printf("Congratulations! you are going to attend Kubecon from %v ", city)
	case "Manchester":
		fmt.Printf("Congratulations! you are going to attend Kubecon from %v ", city)
	default:
		fmt.Printf("You enter invalid city or may be make spelling mistake")
	}
}
