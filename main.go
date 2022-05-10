package main

import (
	"booking-app/helper"
	"fmt" // fmt pkg allows for input/output
	"strconv"
	// strings pkg allows for separation on a space character in a field
)

// DEFINING PACKAGE LEVEL VARS are available to ALL functions on this page
// other variables should be as local as possible -- like name and email and usertickets
const conferenceTickets int = 50
var remainingTickets uint = 50
var conferenceName string = "Go Conference" // OR conferenceName := "Go Conference"
var bookings = make([]map[string]string, 0) // no defined length makes this a LIST OF MAPS



func main() {

	// greetUsers(conferenceName, conferenceTickets,  remainingTickets)
	greetUsers()

		// A PLAIN FOR LOOP that keeps running while there are tickets left
	for {
	// get multi vars data from the same func
		firstName, lastName, email, userTickets := getUserInput()
		// the capV below is receiving an exported value frm th helper package we created
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)


		// // validate user input func
		// validateUserInput(firstName, lastName, email, userTickets)
		
		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := printFirstNames()
			fmt.Printf("The first names %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				break



			// // CONLOG NOTES
			// fmt.Println("\n-------------------------")
			// // printing WHOLE array 
			// fmt.Printf("The WHOLE slice is : %v\n", bookings)
			// fmt.Printf("The FIRST value in the array is : %v\n", bookings[0])
			// fmt.Printf("The slice TYPE is : %T\n", bookings)
			// fmt.Printf("The slice LENGTH is : %v\n\n", len(bookings))

			// fmt.Println("-------------------------")
		

			// call function print first names
			// firstNames()
			// receiving a value from the function
			// firstNames := printFirstNames(bookings)
			// fmt.Printf("\n\nThe first names are:  %v\n\n", firstNames)

			// LOGIC for tickets remaing 
			// exit application if no tickets are left
			// if remainingTickets == 0 {
			// 	// END FOR LOOP and program
			// 	fmt.Println("Our conference is booked out. Come back next year.")
			// 	break
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


func printFirstNames() []string {
	// create a SLICE of first names
	firstNames := []string{}
	// print only first names FOR EACH LOOP
	for _, booking := range bookings {	
		// var names = strings.Fields(booking) // before we converted to a userData MAP
		
		firstNames = append(firstNames, booking["firstName"])

		fmt.Printf("LIST OF BOOKINGS VIA MAP is %v:  \n", bookings)
	}

	// fmt.Printf("\n\nThe first names are:  %v\n\n", firstNames)
	// RETURNING A VALUE also has to be stated at the top of this func between the () and {} and you return the TYPE
	return firstNames
}


func getUserInput() (string, string, string, uint) {	
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

// MOVED TO HELPER.GO 
// func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
// 	isValidName := len(firstName) >= 2 && len(lastName) >= 2
// 	isValidEmail := strings.Contains(email, "@")
// 	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
// 	// can return MULTIPLE values from a function  instead of 1
// 	return isValidName, isValidEmail, isValidTicketNumber
// }

func greetUsers() {
		fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
		fmt.Println("")
	}

func bookTicket(userTickets uint, firstName string, lastName string,  email string ){
	// LOGIC for book ticket in system
	remainingTickets = remainingTickets - userTickets

	// // putting items into the array called bookings
	// bookings = append(bookings, firstName+" "+lastName)

	// create a MAP for a user with datatypes and datavalues
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets),10) //converts uint into string

	bookings = append(bookings, userData)



	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
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