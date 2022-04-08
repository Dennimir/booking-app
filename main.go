package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

const conferenceName string = "Go Conference"
const conferenceTickets uint8 = 50

var bookings = []UserData{} //make([]map[string]string)
var remainingtickets uint8 = conferenceTickets

type UserData struct {
	userFirstName string
	userLastName  string
	userEmail     string
	userTickets   uint8
}

func main() {

	greetUsers()

	for {
		userFirstName, userLastName, userEmail, userTickets := receiveUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.InputValidation(userFirstName, userLastName, userEmail, userTickets, remainingtickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, userFirstName, userLastName, userEmail)
			sendTicket(userTickets, userFirstName, userLastName, userEmail)
			printFirstNames(bookings)

			if remainingtickets == 0 {
				//end program
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			errorPrint(isValidName, isValidEmail, isValidTicketNumber)
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingtickets)
	fmt.Println("Get your ticket here to attend")
}

func receiveUserInput() (string, string, string, uint8) {
	var userFirstName string
	var userLastName string
	var userEmail string
	var userTickets uint8

	fmt.Println("Enter your first name:")
	fmt.Scan(&userFirstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&userLastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&userEmail)

	fmt.Printf("%v, how many tickets do you want to book:\n", userFirstName)
	fmt.Scan(&userTickets)
	return userFirstName, userLastName, userEmail, userTickets
}

func bookTicket(userTickets uint8, userFirstName string, userLastName string, userEmail string) {
	remainingtickets = remainingtickets - userTickets

	//create a map for user
	var userData = UserData{
		userFirstName: userFirstName,
		userLastName:  userLastName,
		userEmail:     userEmail,
		userTickets:   userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive a confirmation email at %v\n",
		userFirstName, userLastName, userTickets, userEmail)
	fmt.Printf("%v tickets remaining for \"%v\"\n", remainingtickets, conferenceName)
}

func printFirstNames(bookings []UserData) {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.userFirstName)
	}
	fmt.Printf("\nThe first names of bookings are: %v\n\n", firstNames)
}

func errorPrint(isValidName bool, isValidEmail bool, isValidTicketNumber bool) {
	if !isValidName {
		fmt.Println("First name or last name you entered is too short")
	}
	if !isValidEmail {
		fmt.Println("Email you entered does not contain @ sign")
	}

	if !isValidTicketNumber {
		fmt.Println("Number of tickets you entered is invalid")
	}
}

func sendTicket(userTickets uint8, userFirstName string, userLastName string, userEmail string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, userFirstName, userLastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email %v\n", ticket, userEmail)
	fmt.Println("#################")
}
