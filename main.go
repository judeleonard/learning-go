package main

import (
	"booking-app/helper"
	"fmt"
)

const ConferenceTickets = 50
var ApplicationName = "Hotel Bookings App"
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName string
	lastName string
	email 	string
	numberOfTickets uint
}


func main() {

	greetUsers()

	for {
		// get user inputs
		firstName, lastName, email, userTickets := getUserInput()

		// validate user inputs
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// logic to make sure the user doesn't book more than the available tickets 
		if isValidName && isValidEmail && isValidTicketNumber {
		
			bookTickets(userTickets, firstName, lastName, email)
			go helper.SendTickets(userTickets, firstName, lastName, email)
			// print first names of users
			firstNames := getFirstNames()
			fmt.Printf("All the first names of the bookings are %v \n", firstNames)

			// logic to check if the available remaining tickets is zero
			// if remaining ticket is zero the program should exit the loop

			if remainingTickets == 0 {
				// end the program
				fmt.Println("Sorry all tickets are sold out come back next week")
				break
			}
			
		} else {
			if !isValidName {
				fmt.Println("The first and last name you entered is too short")

			}
			if !isValidEmail {
				fmt.Println("The email address you entered is incorrect")
			}
			if isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
			//fmt.Printf("We only have %v tickets so you can't book %v tickets \n", remainingTickets, userTickets)

		}
		

	}

	
}

func greetUsers() {
	fmt.Printf("Welcome to our %v \n", ApplicationName)
	fmt.Printf("We have a Total of %v tickets and %v tickets are Available \n", ConferenceTickets, remainingTickets)
	fmt.Println("Good to have you here!")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		//var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
	
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName  string
	var email 	  string
	var userTickets uint

	// ask user for their first name
	println("Enter your first Name:")
	fmt.Scan(&firstName)

	// ask user for their last name
	println("Enter your Last Name:")
	fmt.Scan(&lastName)

	// ask user for their email address
	println("Enter your Email:")
	fmt.Scan(&email)

	println("Enter the number of tickets you want:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create userdata structure
	var userData = userData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v \n", bookings)

	fmt.Printf("The whole Slice %v \n", bookings)
	fmt.Printf("The first value %v \n", bookings[0])
	fmt.Printf("The type of the Slice %T\n", bookings)
	fmt.Printf("The length of the array %v \n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf(" %v tickets is remaining for %v \n", remainingTickets, ConferenceTickets)
}
