package main

import (
	"fmt"
	"sync"
	"time"
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

// ? add waitgroup
var wg = sync.WaitGroup{}

// ? main function
func main() {
	greetUser()
	for remainingTickets > 0 { //? loop until the tickets are sold out
		var firstName, lastName, userEmail string
		var userTickets int

		//? get the user input for booking
		firstName, lastName, userEmail, userTickets = getUserInput()

		//? book the tickets
		if validateUserInputs(firstName, lastName, userEmail, userTickets, remainingTickets) {
			bookTicket(firstName, lastName, userEmail, userTickets)
			wg.Add(1)
			go sendTicket(firstName, lastName, userTickets, userEmail) //? go routine to send the ticket in the background while the user can book more tickets
		}
		wg.Wait()
	}
	fmt.Println("Our tickets are sold out. Come back next year.")
}

//* Functions

// ? Greets the user and displays the remaining tickets
func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and the remaining tickets are %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend..")
}

func getFirstNames() []string {
	var firstNames []string
	for _, value := range bookings {
		firstNames = append(firstNames, value.firstName) //? here value is the struct and we are accessing the firstName field
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var (
		firstName   string
		lastName    string
		userEmail   string
		userTickets int
	)

	prompts := map[string]interface{}{  
		"Enter your first name: ":                        &firstName,
		"Enter your last name: ":                         &lastName,
		"Enter your email: ":                             &userEmail,
		"Enter the number of tickets you want to book: ": &userTickets,
	}

	//? range over the map the key and value are returned
	for prompt, variable := range prompts {
		fmt.Print(prompt)
		fmt.Scan(variable)
		// Consume the newline character
		fmt.Scanln()

	}

	return firstName, lastName, userEmail, userTickets
}

func bookTicket(firstName, lastName, userEmail string, userTickets int) {
	userName := firstName + " " + lastName
	remainingTickets -= userTickets

	// syntax : var a = make(map[KeyType]ValueType)
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


//? to import a function from another package we use import "<modulename>/<packagename>"
//? let's say we import the helper function in helper package then we would use import "booking-app/helper"
//? then to use the function in the pkg we use helper.ValidateUserInput()  
//? like this we can import any files from different package and manage our code.
