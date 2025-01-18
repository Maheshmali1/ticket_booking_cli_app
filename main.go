package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceTickets int = 50
	remainingTickets := conferenceTickets
	ticketBuyers := []string{}
	ticketBoughts := []int{}

	fmt.Println("Welcome to Go programming conference!")
	fmt.Printf("We have total of %v tickets for conference\n", conferenceTickets)

	for {
		fmt.Printf("There are %v tickets remaining, Hurry up!!\n", remainingTickets)
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := email != "" && strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			ticketBuyers = append(ticketBuyers, firstName+" "+lastName)
			ticketBoughts = append(ticketBoughts, userTickets)
			fmt.Printf("You have booked %v tickets\n", userTickets)
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

	fmt.Println("Conference ticket booking is over!")
	fmt.Println("Welcome gophers to conference!")
	for index, value := range ticketBuyers {
		fmt.Printf("Gopher : %v has booked %v tickets...\n", value, ticketBoughts[index])
	}

}
