package main

import (
	"fmt"
	"strings"
)


//Capitalize the func header to export it 
func validateUserInputs(firstName string, lastName string, userEmail string, userTickets int, remainingTickets int) bool {
	if len(firstName) < 2 || len(lastName) < 2 {
		fmt.Println("Please enter a valid name..!")
		return false
	}

	if !strings.Contains(userEmail, "@") || !strings.Contains(userEmail, ".com") {
		fmt.Println("Please enter a valid email..!")
		return false
	}

	if userTickets == 0 {
		fmt.Println("Please enter a valid number..!")
		return false
	}

	if userTickets > remainingTickets {
		fmt.Printf("Sorry, we don't have that many tickets left. Remaining tickets are %v\n", remainingTickets)
		return false
	}
	return true
}



//? for exporting a function capitalize the first letter. 
//? let's say if the helper function was in the helper/helper.go then we would need to export it like this 
//? func ValidateUserInputs() - V capitalized. for exporting.
