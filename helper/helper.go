package helper

import "strings"

func InputValidation(userFirstName string, userLastName string, userEmail string, userTickets uint8, remainingtickets uint8) (bool, bool, bool) {
	isValidName := len(userFirstName) > 1 && len(userLastName) > 1
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingtickets
	return isValidName, isValidEmail, isValidTicketNumber
}
