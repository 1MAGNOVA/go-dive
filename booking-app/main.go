package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 1000
	var remainingTickets uint = 1000

	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("we have total of %v and %v  remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("You can now get a discount when you buy 4 tickets or more")

	var firstName string
	var lastName string
	var email string
	// ask user for name
	bookings := []string{}

	var userTickets uint

	fmt.Println("Enter your first name here:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name here:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email here:")
	fmt.Scan(&email)
	fmt.Println("How many Tickets would you like to buy?:")
	fmt.Scan(&userTickets)
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, (firstName + " " + lastName))

	fmt.Printf("%v, %v with email address: %v, booked %v amount of tickets.\n", firstName, lastName, email, userTickets)
	fmt.Printf("You have %v tickets left\n", remainingTickets)

	fmt.Printf("%v reserved bookings\n", bookings[0])
	fmt.Printf("These people have booked tickets %v\n", bookings)

}
