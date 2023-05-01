package main

import (
	"fmt"
	"strconv"
)

// ? package level variables
const conferenceTickets = 50

var conferenceName = "Go conference"
var remainingTickets = conferenceTickets
var bookings = make([]map[string]string, 0)

// ? main function
func main() {

	greetUser()

	for remainingTickets > 0 {
		var firstName, lastName, userEmail string
		var userTickets int

		//? get the user input for booking
		firstName, lastName, userEmail, userTickets = getUserInput()

		//? book the tickets
		if validateUserInputs(firstName, lastName, userEmail, userTickets, remainingTickets) {
			bookTicket(firstName, lastName, userEmail, userTickets)
		}
	}

	fmt.Println("Our tickets are sold out. Come back next year.")

}

// ? functions
func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and the remaining tickets are %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend..")
}

func getFirstNames() []string {
	var firstNames []string
	for _, value := range bookings {
		firstNames = append(firstNames, value["firstName"])
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
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

	return firstName, lastName, userEmail, userTickets
}

func bookTicket(firstName, lastName, userEmail string, userTickets int) {
	userName := firstName + " " + lastName
	remainingTickets -= userTickets

	//syntax : var a = make(map[KeyType]ValueType)
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = userEmail
	userData["tickets"] = strconv.FormatInt(int64(userTickets), 10)

	//? append the map to the slice

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, userTickets, userEmail)
	fmt.Printf("Remaining tickets are %v\n", remainingTickets)
	fmt.Printf("These are all the bookings: %v\n", getFirstNames())
}
