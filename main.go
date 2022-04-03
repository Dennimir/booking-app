package main

import "fmt"

func main() {
	bookings := []string{}
	var userFirstName string
	var userLastName string
	var userEmail string
	var userTickets uint8
	const conferenceName string = "Go Conference"
	const conferenceTickets uint8 = 50
	remainingtickets := conferenceTickets

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingtickets)
	fmt.Println("Get your ticket here to attend")

	fmt.Println("Enter your first name:")
	fmt.Scan(&userFirstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&userLastName)
	fmt.Println("Enter your email:")
	fmt.Scan(&userEmail)
	fmt.Printf("%v, how many tickets do you want to book:\n", userFirstName)
	fmt.Scan(&userTickets)

	remainingtickets = remainingtickets - userTickets
	bookings = append(bookings, userFirstName+" "+userLastName)

	fmt.Printf("These are all our bookings: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive a confirmation email at %v\n",
		userFirstName, userLastName, userTickets, userEmail)
	fmt.Printf("%v tickets remaining for \"%v\"\n", remainingtickets, conferenceName)

}
