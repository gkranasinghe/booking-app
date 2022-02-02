package main

import "fmt"

func main(){
	var conferenceName = "Go Conference"
	const conferenceTickets = 50 
	var remainingTickets uint = 50 
    
    fmt.Printf("conferenceTickets is %T\n",conferenceTickets)
	fmt.Printf("remainingTickets is %T\n",remainingTickets)

	fmt.Printf("Welcome to %v  \n",conferenceName)
	fmt.Printf("We have a total of  %v tickets and %v remainingTickets availlable " ,conferenceTickets,remainingTickets)
	fmt.Println("Get tickets to atttend here  ")
     

	var userName string
	var userTickets int 

	userName = "Tom"
	userTickets = 2


	fmt.Printf("User %v booked %v tickets.\n" ,userName,userTickets)


}

