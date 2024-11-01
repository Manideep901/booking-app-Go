package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenaceTickets int = 50
	var remainingTickets int = 50
	var bookings []string
	log.Printf("Welcome to %v booking application\n", conferenceName)
	log.Printf("We have total of %v tickets and %v are still available.\n", conferenaceTickets, remainingTickets)
	log.Println("Get your tickets here to attend")
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int
		log.Print("Enter your  first name: ")
		fmt.Scan(&firstName)
		log.Print("Enter your last name: ")
		fmt.Scan(&lastName)
		log.Print("Enter your email: ")
		fmt.Scan(&email)
		log.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)
		remainingTickets -= userTickets
		bookings = append(bookings, firstName+" "+lastName)
		log.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email);
		log.Printf("%v tickets remaining for %v\n", remainingTickets, conferenaceTickets)

		firstNames := [] string{}
		for _,booking := range bookings {
			var names =	strings.Fields(booking);
			firstNames = append(firstNames,names[0])
		}
		log.Printf("The first names of bookings are: %v\n",firstNames)
		 if remainingTickets== 0{
			log.Println("Our conference is booked out. Come back next year.");
			break;
		 }
	}

}
