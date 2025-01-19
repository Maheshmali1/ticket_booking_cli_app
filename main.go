package main

import (
	"fmt"
	"sync"
	"ticket-booking-app/shared"
	"time"
)

const conferenceTickets int = 50

var remainingTickets = conferenceTickets
var ticketBuyers = make([]userData, 0)
var waitGroup = sync.WaitGroup{}

type userData struct {
	firstName   string
	lastName    string
	email       string
	userTickets int
}

func main() {
	greetUsers()

	for {
		fmt.Printf("There are %v tickets remaining, Hurry up!!\n", remainingTickets)
		firstName, lastName, email, userTickets := shared.GetUserInputs()

		isValidName, isValidEmail, isValidTicketNumber := shared.ValidateUserInputs(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(firstName, lastName, email, userTickets)
		} else {

			fmt.Println("Your inputs are invalid.. Please try again...")
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
	waitGroup.Wait()
	conferenceBookingEndingResponse()
}

func greetUsers() {
	fmt.Println("Welcome to Go programming conference!")
	fmt.Printf("We have total of %v tickets for conference\n", conferenceTickets)
}

func bookTickets(firstName string, lastName string, email string, userTickets int) {
	remainingTickets = remainingTickets - userTickets
	user := userData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}
	ticketBuyers = append(ticketBuyers, user)
	waitGroup.Add(1)
	go sendTicketsToUser(user)
	fmt.Printf("Thank you %v for booking %v tickets.\n", firstName+" "+lastName, userTickets)
}

func conferenceBookingEndingResponse() {
	fmt.Println("******* Conference ticket booking is over! ******* ")
	fmt.Println("🚀🚀 Welcome gophers to conference!")
	for _, value := range ticketBuyers {
		fmt.Printf("Gopher 👨‍💻: %v has booked %v tickets...\n", value.firstName+" "+value.lastName, value.userTickets)
	}
}

func sendTicketsToUser(user userData) {
	emailBody := fmt.Sprintf("Hello %v %v, \nYou have booked %v tickets for conference.\nYou can contact us at: abc@example.com\n", user.firstName, user.lastName, user.userTickets)

	fmt.Println("Sending ticket:\n", emailBody)
	time.Sleep(10 * time.Second)
	fmt.Println("Ticket has been sent to", user.email)
	waitGroup.Done()
}
