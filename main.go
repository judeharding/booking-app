package main

// the fmt package is used for input/output functionality
import "fmt"

func main()   {

	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	// array that can hold 50 items
	var bookings [50]string

	// a flexible array is called a slice -- use a slice for better memory management
	var bookingslice [] string

	// different syntax for same thing above but ONLY works on VARS
	//  conferenceName  := "Go Conference"
	//  remainingTickets  := 50 // can use on int but NOT uint


	// finding out what data type the vars below are 
	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)


	fmt.Println("Welcome to our ", conferenceName, " booking application")
	// fmt.Println("We have a total of ", conferenceTickets, " and ", remainingTickets, " are still available")
	// Printf allows you to print the %VARIABLES in different formats - see documentation 
	// below uses place holders with printF and above is the standard way with printLN 
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")



	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println(remainingTickets)
	fmt.Println(&remainingTickets)



	// we cannot assign "tom" to the UN but we can POINT (&) to the memory address where the user enters the name AND
	// the program will wait until you enter a name in terminal 
	// userName = "Tom"
	fmt.Println("Enter your first name:  ")
	fmt.Scan(&firstName)
  // SCAN pauses the program and waits on user entry
	// Pointers & are needed when waiting on data to be entered

	fmt.Println("Enter your last name:  ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email:  ")
	fmt.Scan(&email)
	fmt.Println("How many tickets do you want:  ")
	fmt.Scan(&userTickets)

  remainingTickets = remainingTickets - userTickets
  bookings[0] = firstName  + " " + lastName
	// no need for an index when using SLICE instead of ARRAY 
	bookingslice = append(bookingslice, firstName + " " + lastName)

	fmt.Println("-----")
	fmt.Printf("Thank you %v %v for purchasing %v tickets.  You will receive a confirmation email at %v . \n", firstName, lastName, userTickets, email)

	fmt.Println("-----")
	fmt.Printf("%v tickets remain open for t %v \n", remainingTickets, conferenceName)

	fmt.Println("-----")
	fmt.Printf("The WHOLE array%v \n", bookings)
	fmt.Printf("The first value in the array%v \n", bookings[0])
	fmt.Printf("The array type is %T \n", bookings)
	fmt.Printf("The array length is %v \n", len(bookings))

	fmt.Println("-----")
	fmt.Printf("The WHOLE SLICE ARRAY is %v \n", bookingslice)
	fmt.Printf("The first value in the SLICE%v \n", bookingslice[0])
	fmt.Printf("The SLICE type is %T \n", bookingslice)
	fmt.Printf("The SLICE length is %v \n", len(bookingslice))
}

