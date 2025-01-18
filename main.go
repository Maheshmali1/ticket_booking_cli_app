package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var remainingTickets = conferenceTickets
var ticketBuyers = []string{}
var ticketBoughts = []int{}

func main() {
	greetUsers()

	for {
		fmt.Printf("There are %v tickets remaining, Hurry up!!\n", remainingTickets)
		firstName, lastName, email, userTickets := getUserInputs()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInputs(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(firstName, lastName, userTickets)
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("Invalid email address")
			}

			if !isValidTicketNumber {
				fmt.Println("Invalid number of tickets")
			}
		}

		fmt.Println("Do you want to continue? (yes/no)")
		var continueBooking string
		fmt.Scan(&continueBooking)
		if continueBooking != "yes" {
			break
		}
	}

	conferenceBookingEndingResponse()

}

func greetUsers() {
	fmt.Println("Welcome to Go programming conference!")
	fmt.Printf("We have total of %v tickets for conference\n", conferenceTickets)
}

func getUserInputs() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets you want to book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func validateUserInputs(firstName string, lastName string, email string, userTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := email != "" && strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func bookTickets(firstName string, lastName string, userTickets int) {
	remainingTickets = remainingTickets - userTickets
	ticketBuyers = append(ticketBuyers, firstName+" "+lastName)
	ticketBoughts = append(ticketBoughts, userTickets)
	fmt.Printf("Thank you %v for booking %v tickets.\n", firstName+" "+lastName, userTickets)
}

func conferenceBookingEndingResponse() {
	fmt.Println("Conference ticket booking is over!")
	fmt.Println("Welcome gophers to conference!")
	for index, value := range ticketBuyers {
		fmt.Printf("Gopher : %v has booked %v tickets...\n", value, ticketBoughts[index])
	}
}
