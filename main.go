package main

// the fmt package is used for input/output functionality
import "fmt"

func main()   {

	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	// different syntax for same thing above but ONLY works on VARS
	//  conferenceName  := "Go Conference"
	//  remainingTickets  := 50 // can use on int but NOT uint


	// finding out what data type the vars below are 
	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)


	fmt.Println("Welcome to our ", conferenceName, " booking application")
	// fmt.Println("We have a total of ", conferenceTickets, " and ", remainingTickets, " are still available")
	// Printf allows you to print the %VARIABLES in different formats - see documentation 
	// below uses place holders and above is the standard way 
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")



	var userName string
	var userTickets int

	fmt.Println(remainingTickets)
	fmt.Println(&remainingTickets)



	// we cannot assign "tom" to the UN but we can POINT (&) to the memory address where the user enters the name AND
	// the program will wait until you enter a name in terminal 
	// userName = "Tom"
	fmt.Println("Enter your first name:  ")
	fmt.Scan(&userName)


	userTickets = 2

	fmt.Println("-----")
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}

