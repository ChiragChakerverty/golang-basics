package main

import (
	"fmt"
)

func main(){

	ConferenceName := "Ted Talk"
	const numberOfTickets = 50
	remainingTickets := 50
	var bookings = make([]map[string]string , 0)

	greetUsers(ConferenceName , numberOfTickets , remainingTickets)

	for	{
		FirstName , LastName , email , UserTickets := takingUserInput()

		isNameValid , isEmailValid , isUserTicketValid , userInputValid := isUserInputValid(FirstName , LastName , email , &UserTickets , &remainingTickets)
		
		if (userInputValid) {
			bookings = PrintBookingDetails(FirstName , LastName , email , &UserTickets , &remainingTickets , bookings)
		} else {
			PrintReasonOfInvalidation(isNameValid , isEmailValid , isUserTicketValid)
		}
	
		fmt.Println("All Userdata of the bookings done is :" , bookings)

		if remainingTickets <=0 {
			break
	 	}
	}

	FirstNames := giveFirstNames(bookings)
	fmt.Println("All the firstnames of the people who booked the tickets are are :" , FirstNames)
}




