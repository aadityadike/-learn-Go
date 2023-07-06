package otherconcepts

import (
	"errors"
	"fmt"
)

func Errors() {
	variable, err := getUser()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}
	fmt.Printf("User: %v", variable)
}

func getUser() (string, error){
	user := fmt.Sprint("aditya dike")
	return user, errors.New("something is wrong")
}
