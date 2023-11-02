package helper

import (
	"strings"
	"fmt"
	"time"
)

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber

}

func SendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(40 * time.Second)
	var tickets = fmt.Sprintf("%v Tickets for %v %v \n", userTickets, firstName, lastName)
	fmt.Println("########################")
	fmt.Printf("Sending tickets: \n %v \n to email address %v \n", tickets, email)
	fmt.Println("########################")

}