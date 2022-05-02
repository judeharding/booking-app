package main

import (
	"fmt"     // fmt pkg allows for input/output
	"strings" // strings pkg allows for separation on a space character in a field
)

func main() {

	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var conferenceName string = "Go Conference" // OR conferenceName := "Go Conference"
	var bookings = []string{} // no defined length makes this a SLICE

	fmt.Printf("\n\nWelcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)


		// A PLAIN FOR LOOP that keeps running while there are tickets left
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// asking for user input
		fmt.Println("*****  NEW RECORD  *****")
		fmt.Println("Enter Your First Name: ")
		fmt.Scanln(&firstName)

		fmt.Println("Enter Your Last Name: ")
		fmt.Scanln(&lastName)

		fmt.Println("Enter Your Email: ")
		fmt.Scanln(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scanln(&userTickets)

		// validate user input
		var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		var isValidEmail bool = strings.Contains(email, "@")
		var isValidTicketNumber bool = userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {

			// LOGIC for book ticket in system
			remainingTickets = remainingTickets - userTickets

			// putting items into the array called bookings
			bookings = append(bookings, firstName+" "+lastName)


			// CONLOG NOTES
			fmt.Println("\n-------------------------")
			// printing WHOLE array 
			fmt.Printf("The WHOLE slice is : %v\n", bookings)
			fmt.Printf("The FIRST value in the array is : %v\n", bookings[0])
			fmt.Printf("The slice TYPE is : %T\n", bookings)
			fmt.Printf("The slice LENGTH is : %v\n\n", len(bookings))

			fmt.Println("-------------------------")
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n\n", remainingTickets, conferenceName)
			fmt.Println("*****  END OF RECORD  *****")


			// create a SLICE of first names
			firstNames := []string{}
			// print only first names FOR EACH LOOP
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names %v\n", firstNames)

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