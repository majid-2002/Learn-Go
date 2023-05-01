package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go conference"
	const conferenceTickets = 50
	remainingTickets := conferenceTickets
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and the remaining tickets are %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend..")

	for remainingTickets > 0 {
		var firstName, lastName, userEmail string
		var userTickets int

		fmt.Print("\nEnter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email: ")
		fmt.Scan(&userEmail)

		fmt.Print("Enter the number of tickets you want to book: ")
		fmt.Scan(&userTickets)

		if len(firstName) < 2 || len(lastName) < 2 {
			fmt.Println("Please enter a valid name..!")
			continue
		}

		if !strings.Contains(userEmail, "@") || !strings.Contains(userEmail, ".com") {
			fmt.Println("Please enter a valid email..!")
			continue
		}

		if userTickets == 0 {
			fmt.Println("Please enter a valid number..!")
			continue
		}

		if userTickets > remainingTickets {
			fmt.Printf("Sorry, we don't have that many tickets left. Remaining tickets are %v\n", remainingTickets)
			continue
		}

		userName := firstName + " " + lastName
		remainingTickets -= userTickets
		bookings = append(bookings, userName)

		fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, userTickets, userEmail)
		fmt.Printf("Remaining tickets are %v\n", remainingTickets)

		var firstNames []string
		for _, value := range bookings {
			firstNames = append(firstNames, strings.Fields(value)[0])
		}
		fmt.Printf("These are all the bookings: %v\n", firstNames)
	}

	fmt.Println("Our tickets are sold out. Come back next year.")
}
