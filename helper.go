/**
 * This is another file with classes defined in the same package as the main.go file.
 * As an alternative, we could have created a separate dir to put this file into, give it a different package name
 * and let it get called through import using "module_name/package" format. Ex: golang-studies/helper
 */
package main

import (
	"fmt"
	"strings"
	"time"
)

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	// User Input validation
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketsNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketsNumber
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// Using Sleep function to simulate data being transmitted
	time.Sleep(10 * time.Second)

	var tickets = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#######################")
	fmt.Printf("Sending ticket:\n%v\nto email address %v.\n", tickets, email)
	fmt.Println("#######################")

	wg.Done()
}
