package main

import "fmt"

func main(){
	var conferenceName = "Go Conference"
	const conferenceTickets = 50 
	var remainingTickets uint = 50 
    
    fmt.Printf("conferenceTickets is %T\n",conferenceTickets)
	fmt.Printf("remainingTickets is %T\n",remainingTickets)

	fmt.Printf("Welcome to %v  \n",&conferenceName)
	fmt.Printf("We have a total of  %v tickets and %v remainingTickets availlable " ,conferenceTickets,remainingTickets)
	fmt.Println("Get tickets to atttend here  ")
     

	var firstName string
	var lastName string
	var email string
	var userTickets int 


	fmt.Println("Enter first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter email:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	fmt.Printf("User %v %v  booked %v tickets\nyou will recieve confirmation at %v \n" ,firstName,lastName,userTickets,email)


}

