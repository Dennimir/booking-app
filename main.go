package main

import "fmt"

func main() {
	const conferenceName string = "Go Conference"
	const conferenceTickets uint8 = 50
	remainingtickets := conferenceTickets

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingtickets)
	fmt.Println("Get your ticket here to attend")

	var userName string
	var userTickets int
	//ask user for the name
	userName = "Tom"
	userTickets = 1
	fmt.Print(userName, userTickets)

}
