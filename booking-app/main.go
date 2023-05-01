package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go conference"
	const conferenceTickets int = 50
	remainingTickets := 50    // syntatic sugar
	var bookings = []string{} //? this is a slice. A slice is just like an array but like a dynamic size array

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and the remaining tickets are", remainingTickets)
	fmt.Println("Get your tickets to attend..")

	for {

		var firstName string
		var lastName string
		var userTickets int
		var userEmail string

		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email")
		fmt.Scan(&userEmail)

		fmt.Println("Enter the number of tickets you want to book")
		fmt.Scan(&userTickets)

		var userName string = firstName + " " + lastName

		remainingTickets -= userTickets

		//? this is how we add a new item to the slice by using the append method in go lang
		bookings = append(bookings, userName)

		// fmt.Printf("the whole slice : %v\n", bookings)
		// fmt.Printf("the first valu : %v\n", bookings[0])
		// fmt.Printf("Slice type : %T\n", bookings)
		// fmt.Printf("Slice length : %v\n", len(bookings))

		fmt.Printf("Thank you %v for booking %v tickets. You will recieve a confirmation email at %v\n", userName, userTickets, userEmail)
		fmt.Println("Remaining tickets are", remainingTickets)

		var firstNames = []string{}

		for _, value := range bookings {
			var names = strings.Fields(value)
			firstNames = append(firstNames, names[0])
		}

		fmt.Println("These are all the bookings", firstNames)
	}

}
