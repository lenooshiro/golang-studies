package main

import (
	"fmt"
	"sync"
)

// You can use this sintax, but can't use with const nor explicitly declare a data type
// conferenceName := "Go Conference"

// These are "package level" variables
const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50

// This is an array. Arrays always have fixed size
// var bookings [50]string
// This is a slice
// var bookings []string
// A slice can also be defined like this
// bookings := []string{}

// This is a list of maps
// var bookings = make([]map[string]string, 0)

// This is a list of a struct
var bookings = make([]UserData, 0)

// This is a struct data type
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// A struct used to add goroutines (aka subthreads) so the main thread can know
// how many it needs to wait before taking some action.
var wg = sync.WaitGroup{}

func main() {
	// Calling functions with arguments
	greetUsers()

	// For loop sintax
	// for {

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketsNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	// Using if-else sintax
	if isValidName && isValidEmail && isValidTicketsNumber {
		bookTicket(userTickets, firstName, lastName, email)

		// This will wait for the subthread to finish
		wg.Add(1)
		// The "go" keyword starts the following call in a separate thread to solve Concurrency
		go sendTicket(userTickets, firstName, lastName, email)

		fmt.Printf("These are all our bookings first names: %v\n", getFirstNames())

		if remainingTickets == 0 {
			fmt.Printf("Our conference is booked out. Come back next year!")
			//break
		}
	} else {
		if !isValidName {
			fmt.Printf("First name or last name you entered is too short!\n")
		}
		if !isValidEmail {
			fmt.Printf("Email address you entered doesn't contain a '@' sign.\n")
		}
		if !isValidTicketsNumber {
			fmt.Printf("Number of tickets you entered is invalid.\n")
		}
	}

	// Switch case sintax example
	city := "London"
	switch city {
	case "New York":
		// some code here for New York
	case "Singapore":
		// some code here for Singapore
	case "Sao Paulo", "Rio de Janeiro":
		// some code here for Sao Paulo and Rio de Janeiro cities
	default:
		// some default code here
	}
	// }

	// Wait for the goroutines/subthreads inside the WaitGroup to finish
	wg.Wait()
}

// How to declare functions
func greetUsers() {
	// Different types of print function
	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// firstNames = append(firstNames, names[0])

		//firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Waiting for user input and saving in a variable
	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// Create a Map for the user
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// Create a struct for UserData
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// Add the user Map to our list of Maps
	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v.\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v .\n", firstName, lastName, userTickets, email)
	fmt.Printf("Remaining tickets for %v: %v\n", conferenceName, remainingTickets)
}
