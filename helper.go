package main

import (
	"fmt"
	"strings"
)

func validateUserInput(firstName string, lastName string, emailId string, city string, numberOfSeats uint, remaingSeats uint) bool {

	isUserInputValid := true

	if len(firstName) < 3 {
		fmt.Printf("First name entered %v is less than %v character\n", firstName, 3)
		isUserInputValid = false
	}
	if len(lastName) < 3 {
		fmt.Printf("Last name entered %v is less than %v character\n", lastName, 3)
		isUserInputValid = false
	}
	if len(emailId) < 3 {
		fmt.Printf("First name entered %v is less than %v character\n", emailId, 3)
		isUserInputValid = false
	}

	if strings.Contains(emailId, "@") {
		fmt.Printf("Invalid Email format. Please enter X@X.X format\n")
		isUserInputValid = false
	}

	if len(city) != 1 && (city != "C" || city != "B" || city != "H") {
		fmt.Printf("City name entered %v is not matching with the list provided %v character\n", city, "C/B/H")
		isUserInputValid = false
	}

	if numberOfSeats > remaingSeats {
		fmt.Printf("You entered more than available seats %v. Please re-enter availabel seats.\n", remaingSeats)
		isUserInputValid = false
	}

	return isUserInputValid
}
