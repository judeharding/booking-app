// package main
package helper

import "strings"

//to EXPORT this package, just cap the first letter of the func
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
// func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	// can return MULTIPLE values from a function  instead of 1
	return isValidName, isValidEmail, isValidTicketNumber
}