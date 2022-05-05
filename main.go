package main

import (
	"fmt"     // fmt pkg allows for input/output
	"strings" // strings pkg allows for separation on a space character in a field
)

// DEFINING PACKAGE LEVEL VARS are available to ALL functions on this page
const conferenceTickets uint = 50
var remainingTickets uint = 50
var conferenceName string = "Go Conference" // OR conferenceName := "Go Conference"
var bookings = []string{} // no defined length makes this a SLICE

func main() {

	// greetUsers(conferenceName, conferenceTickets,  remainingTickets)
	greetUsers()

		// A PLAIN FOR LOOP that keeps running while there are tickets left
	for {
	
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		


		// validate user input func
		validateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(remainingTickets, userTickets, firstName, lastName, email)


			// CONLOG NOTES
			fmt.Println("\n-------------------------")
			// printing WHOLE array 
			fmt.Printf("The WHOLE slice is : %v\n", bookings)
			fmt.Printf("The FIRST value in the array is : %v\n", bookings[0])
			fmt.Printf("The slice TYPE is : %T\n", bookings)
			fmt.Printf("The slice LENGTH is : %v\n\n", len(bookings))

			fmt.Println("-------------------------")
		

			// call function print first names
			getFirstNames()
			// receiving a value from the function
			firstNames := getFirstNames(bookings)
			fmt.Printf("\n\nThe first names are:  %v\n\n", firstNames)

			// LOGIC for tickets remaing 
			// exit application if no tickets are left
			if remainingTickets == 0 {
				// END FOR LOOP and program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			// LOGIC for user data entry
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			continue  // continue starts the outside for loop again
		}
	}
}



func greetUsers(){  
	fmt.Printf("\n\nWELCOME NEWBIE to the %v booking app\n", conferenceName)
	fmt.Printf("We have total of  %v tickets and %v are available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Println("")
}

func getFirstNames() []string {
	// create a SLICE of first names
	firstNames := []string{}
	// print only first names FOR EACH LOOP
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	// fmt.Printf("\n\nThe first names are:  %v\n\n", firstNames)
	// RETURNING A VALUE also has to be stated at the top of this func between the () and {} and you return the TYPE
	return firstNames
}


func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint )(bool, bool, bool ){
		isValidName  := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail  := strings.Contains(email, "@")
		isValidTicketNumber  := userTickets > 0 && userTickets <= remainingTickets
		// can return MULTIPLE values from a function  instead of 1
		return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput()(string, string, string, uint){
		// asking for user input
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("*****  NEW RECORD  *****")
		fmt.Println("\n\nEnter Your First Name: ")
		fmt.Scanln(&firstName)

		fmt.Println("Enter Your Last Name: ")
		fmt.Scanln(&lastName)

		fmt.Println("Enter Your Email: ")
		fmt.Scanln(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scanln(&userTickets)

		return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName, string, email string, conferenceName string  ){
	// LOGIC for book ticket in system
	remainingTickets = remainingTickets - userTickets

	// putting items into the array called bookings
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n\n", remainingTickets, conferenceName)
	fmt.Println("*****  END OF RECORD  *****")
}




// 	// Switch statement example
// 	city := "London"
// 	switch city {
// 	case "New York":
// 		// booking for New York conference
// 	case "Singapore", "Hong Kong":
// 		// booking for Singapore & Hong Kong conference
// 	case "London", "Berlin":
// 		// booking for London & Berlin conference
// 	case "Mexico City":
// 		// booking for Mexico City conference
// 	default:
// 		fmt.Print("No valid city selected")
// 	}
// }