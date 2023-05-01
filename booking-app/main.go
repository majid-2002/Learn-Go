package main

import (
	"fmt"
	"time"
	"sync"
)

// ? package level variables
const conferenceTickets = 50

var conferenceName = "Go conference"
var remainingTickets = conferenceTickets
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   int
}

//? add waitgroup
var wg =  sync.WaitGroup{}



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
			wg.Add(1)
			go sendTicket(firstName, lastName, userTickets, userEmail)
		}
		wg.Wait()
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
		firstNames = append(firstNames, value.firstName)
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
	var userData = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     userEmail,
		tickets:   userTickets,
	}

	//? append the struct to the slice
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, userTickets, userEmail)
	fmt.Printf("Remaining tickets are %v\n", remainingTickets)
	fmt.Printf("These are all the bookings: %v\n", getFirstNames())
}

func sendTicket(firstName string, lastName string, userTickets int, userEmail string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket: %v \nto email address %v\n", ticket, userEmail)
	fmt.Println("###############")
	wg.Done()
}
