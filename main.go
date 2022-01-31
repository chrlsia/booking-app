package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to our %v booking system\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTickerNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTickerNumber {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remained for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of booking are: %v\n", firstNames)

			// noTicketsRemaining:=remainingTickets ==0
			if remainingTickets == 0 {
				fmt.Println("Our conference booked out. Come again next year")
				break
			}

		} else {
			if !isValidName{
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail{
				fmt.Println("email address you entered doesn't contain a @ sign")
			}
			if !isValidTickerNumber{
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
}
