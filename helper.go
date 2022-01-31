package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickerNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTickerNumber
}

//we have to import strings
//and now we give the command: go run main.go helper.go
//or go run .